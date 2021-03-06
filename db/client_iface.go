package db

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	bolt "github.com/coreos/bbolt"
	"github.com/golang/protobuf/proto"
	"github.com/iancoleman/strcase"
	"github.com/jinzhu/inflection"
	log "github.com/sirupsen/logrus"

	"jaytaylor.com/andromeda/domain"
	"jaytaylor.com/andromeda/pkg/contains"
)

const (
	TableMetadata          = "andromeda-metadata"
	TablePackages          = "packages"
	TablePendingReferences = "pending-references"
	TableCrawlResults      = "crawl-result"
	TableToCrawl           = "to-crawl"

	MaxPriority = 10 // Number of supported priorities, 1-indexed.
)

var (
	ErrKeyNotFound                = errors.New("requested key not found")
	ErrNotImplemented             = errors.New("function not implemented")
	ErrMetadataUnsupportedSrcType = errors.New("unsupported src type: must be an []byte, string, or proto.Message")
	ErrMetadataUnsupportedDstType = errors.New("unsupported dst type: must be an *[]byte, *string, or proto.Message")

	DefaultQueuePriority     = 3
	DefaultBoltQueueFilename = "queue.bolt"

	tables = []string{
		TableMetadata,
		TablePackages,
		TableToCrawl,
		TablePendingReferences,
		TableCrawlResults,
	}

	qTables = []string{
		TableToCrawl,
		TableCrawlResults,
	}

	// pkgSepB is a byte array of the package component separator character.
	// It's used for hierarchical searches and lookups.
	pkgSepB = []byte{'/'}
)

type Client interface {
	Open() error                                                                                   // Open / start DB client connection.
	Close() error                                                                                  // Close / shutdown the DB client connection.
	Destroy(tables ...string) error                                                                // Destroy K/V tables and / or queue topics.
	EachRow(table string, fn func(k []byte, v []byte)) error                                       // Invoke a callback on the key/value pair for each row of the named table.
	EachRowWithBreak(table string, fn func(k []byte, v []byte) bool) error                         // Invoke a callback on the key/value pair for each row of the named table until cb returns false.
	PackageSave(pkgs ...*domain.Package) error                                                     // Performs an upsert merge operation on a fully crawled package.
	PackageDelete(pkgPaths ...string) error                                                        // Delete a package from the index.  Complete erasure.
	Package(pkgPath string) (*domain.Package, error)                                               // Retrieve a specific package..
	Packages(pkgPaths ...string) (map[string]*domain.Package, error)                               // Retrieve several packages.
	EachPackage(func(pkg *domain.Package)) error                                                   // Iterates over all indexed packages and invokes callback on each.
	EachPackageWithBreak(func(pkg *domain.Package) bool) error                                     // Iterates over packages until callback returns false.
	PathPrefixSearch(prefix string) (map[string]*domain.Package, error)                            // Search for packages with paths matching a specific prefix.
	PackagesLen() (int, error)                                                                     // Number of packages in index.
	RecordImportedBy(refPkg *domain.Package, resources map[string]*domain.PackageReferences) error // Save imported-by relationship updates.
	CrawlResultAdd(cr *domain.CrawlResult, opts *QueueOptions) error                               // Append a crawl-result to the queue for later merging and save.
	CrawlResultDequeue() (*domain.CrawlResult, error)                                              // Pop a crawl-result from the queue.
	EachCrawlResult(func(cr *domain.CrawlResult)) error                                            // Iterates over all crawl-results and invokes callback on each.
	EachCrawlResultWithBreak(func(cr *domain.CrawlResult) bool) error                              // Iterates over all crawl-results and invokes callback until callback returns false.
	CrawlResultsLen() (int, error)                                                                 // Number of unprocessed crawl results.
	ToCrawlAdd(entries []*domain.ToCrawlEntry, opts *QueueOptions) (int, error)                    // Only adds entries which don't already exist.  Returns number of new items added.
	ToCrawlRemove(pkgs []string) (int, error)                                                      // Scrubs items from queue.
	ToCrawlDequeue() (*domain.ToCrawlEntry, error)                                                 // Pop an entry from the crawl queue.
	EachToCrawl(func(entry *domain.ToCrawlEntry)) error                                            // Iterates over all to-crawl entries and invokes callback on each.
	EachToCrawlWithBreak(func(entry *domain.ToCrawlEntry) bool) error                              // Iterates over to-crawl entries until callback returns false.
	ToCrawlsLen() (int, error)                                                                     // Number of packages currently awaiting crawl.
	MetaSave(key string, src interface{}) error                                                    // Store metadata key/value.  NB: src must be one of raw []byte, string, or proto.Message struct.
	MetaDelete(key string) error                                                                   // Delete a metadata key.
	Meta(key string, dst interface{}) error                                                        // Retrieve metadata key and populate into dst.  NB: dst must be one of *[]byte, *string, or proto.Message struct.
	PendingReferences(pkgPathPrefix string) ([]*domain.PendingReferences, error)                   // Retrieve pending references listing for a package path prefix.
	PendingReferencesSave(pendingRefs ...*domain.PendingReferences) error                          // Save pending references.
	PendingReferencesDelete(keys ...string) error                                                  // Delete pending references keys.
	EachPendingReferences(fn func(pendingRefs *domain.PendingReferences)) error                    // Iterate over each *domain.PrendingReferences object from the pending-references table.
	EachPendingReferencesWithBreak(fn func(pendingRefs *domain.PendingReferences) bool) error      // Iterate over each *domain.PrendingReferences object from the pending-references table until callback returns false.
	PendingReferencesLen() (int, error)                                                            // Number of pending references keys.
	RebuildTo(otherClient Client, kvFilters ...KeyValueFilterFunc) error                           // Rebuild a fresh copy of the DB at destination.  Return ErrNotImplmented if not supported.  Optionally pass in one or more KeyValueFilterFunc functions.
	Backend() Backend                                                                              // Expose underlying backend impl.
	Queue() Queue                                                                                  // Expose underlying queue impl.
}

type KeyValueFilterFunc func(table []byte, key []byte, value []byte) (keyOut []byte, valueOut []byte)

// skipKVFilterFunc is an internal reference used by SkipKVFilter.
func skipKVFilterFunc(table []byte, key []byte, value []byte) (keyOut []byte, valueOut []byte) {
	panic("software author error: should never be invoked")
	return key, value
}

// SkipKVFilter is a signal filter sentinel value to skip Key-Value tables.
var SkipKVFilter = skipKVFilterFunc

// skipQFilterFunc is an internal reference used by SkipQFilter.
func skipQFilterFunc(table []byte, key []byte, value []byte) (keyOut []byte, valueOut []byte) {
	panic("software author error: should never be invoked")
	return key, value
}

// SkipQFilter is a signal filter sentinel value to skip Queue tables.
var SkipQFilter = skipQFilterFunc

type Config interface {
	Type() Type // Configuration type specifier.
}

func NewConfig(driver string, dbFile string) Config {
	switch driver {
	case "bolt", "boltdb":
		return NewBoltConfig(dbFile)

	case "rocks", "rocksdb":
		return NewRocksConfig(dbFile)

	case "postgres", "postgresql", "pg":
		return NewPostgresConfig(dbFile)

	default:
		panic(fmt.Sprintf("Unrecognized or unsupported DB driver %q", driver))
	}
}

// NewClient constructs a new DB client based on the passed configuration.
func NewClient(config Config) Client {
	typ := config.Type()

	switch typ {
	case Bolt:
		be := NewBoltBackend(config.(*BoltConfig))
		// TODO: Return an error instead of panicking.
		if err := be.Open(); err != nil {
			panic(fmt.Errorf("Opening bolt backend: %s", err))
		}
		q := NewBoltQueue(be.db)
		return newKVClient(be, q)

	case Rocks:
		// MORE TEMPORARY UGLINESS TO MAKE IT WORK FOR NOW:
		if err := os.MkdirAll(config.(*RocksConfig).Dir, os.FileMode(int(0700))); err != nil {
			panic(fmt.Errorf("Creating rocks directory %q: %s", config.(*RocksConfig).Dir, err))
		}
		be := NewRocksBackend(config.(*RocksConfig))
		queueFile := filepath.Join(config.(*RocksConfig).Dir, DefaultBoltQueueFilename)
		db, err := bolt.Open(queueFile, 0600, NewBoltConfig("").BoltOptions)
		if err != nil {
			panic(fmt.Errorf("Creating bolt queue: %s", err))
		}
		q := NewBoltQueue(db)
		return newKVClient(be, q)

	case Postgres:
		be := NewPostgresBackend(config.(*PostgresConfig))
		q := NewPostgresQueue(config.(*PostgresConfig))
		return newKVClient(be, q)

	default:
		panic(fmt.Errorf("no client constructor available for db configuration type: %v", typ))
	}
}

type QueueOptions struct {
	Priority        int
	OnlyIfNotExists bool // Only enqueue items which don't already exist.
}

func NewQueueOptions() *QueueOptions {
	opts := &QueueOptions{
		Priority: DefaultQueuePriority,
	}
	return opts
}

// WithClient is a convenience utility which handles DB client construction,
// open, and close..
func WithClient(config Config, fn func(client Client) error) (err error) {
	client := NewClient(config)

	if err = client.Open(); err != nil {
		err = fmt.Errorf("opening DB client %T: %s", client, err)
		return
	}
	defer func() {
		if closeErr := client.Close(); closeErr != nil {
			if err == nil {
				err = fmt.Errorf("closing DB client %T: %s", client, closeErr)
			} else {
				log.Errorf("Existing error before attempt to close DB client %T: %s", client, err)
				log.Errorf("Also encountered problem closing DB client %T: %s", client, closeErr)
			}
		}
	}()

	if err = fn(client); err != nil {
		return
	}

	return
}

// kvTables returns the names of the "regular" key-value tables.
func kvTables() []string {
	kv := []string{}
	for _, table := range tables {
		regular := true
		for _, qTable := range qTables {
			if table == qTable {
				regular = false
				break
			}
		}
		if regular {
			kv = append(kv, table)
		}
	}
	return kv
}

// KVTables publicly exported version of kvTables.
func KVTables() []string { return kvTables() }

// IsKV returns true when s is the name of a Key-Value oriented table.
//
// Note: Does not normalize postgres_formatted_names..
func IsKV(s string) bool {
	return contains.String(kvTables(), s)
}

// QTables returns slice of queue table names.
func QTables() []string {
	tables := []string{}
	for _, table := range qTables {
		tables = append(tables, table)
	}
	return tables
}

// Returns true when the name corresponds with a table.
//
// Note: Does not normalize postgres_formatted_names..
func IsQ(s string) bool {
	return contains.String(qTables, s)
}

// Tables returns a slice of all tables.
func Tables() []string {
	out := []string{}
	for _, table := range tables {
		out = append(out, table)
	}
	return out
}

// StructFor takes a table or queue name and returns a pointer to the newly
// allocated struct of the corresponding type associated with the table.
func StructFor(tableOrQueue string) (proto.Message, error) {
	tableOrQueue = strcase.ToKebab(tableOrQueue)

	if tableOrQueue == "pkg" || tableOrQueue == "pkgs" {
		tableOrQueue = TablePackages
	}

	switch tableOrQueue {
	// N.B.: Metadata type is arbitrary on a per-key basis, so unsupported here.
	// case TableMetadata:

	case inflection.Plural(TablePackages), inflection.Singular(TablePackages):
		return &domain.Package{}, nil

	case inflection.Plural(TablePendingReferences), inflection.Singular(TablePendingReferences):
		return &domain.PendingReferences{}, nil

	case inflection.Plural(TableCrawlResults), inflection.Singular(TableCrawlResults):
		return &domain.CrawlResult{}, nil

	case inflection.Plural(TableToCrawl), inflection.Singular(TableToCrawl):
		return &domain.ToCrawlEntry{}, nil

	default:
		return nil, fmt.Errorf("unrecognized or unsupported table or queue %q", tableOrQueue)
	}
}

// FuzzyTableResolver attempts to resolve the input string to a corresponding table or
// queue name.
//
// An empty string is returned if no match is found.
func FuzzyTableResolver(tableOrQueue string) string {
	if tableOrQueue == "pkg" || tableOrQueue == "pkgs" {
		return TablePackages
	}
	if tableOrQueue == "pending" {
		return TablePendingReferences
	}
	if tableOrQueue == "metadata" || tableOrQueue == "meta" {
		return TableMetadata
	}
	for _, name := range tables {
		if inflection.Singular(name) == tableOrQueue || inflection.Plural(name) == tableOrQueue {
			return name
		}
	}
	return ""
}
