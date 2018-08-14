package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"sync"

	pq "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

var (
	DefaultPostgresConnString = "dbname=andromeda host=/var/run/postgresql" // ssl-mode=verify-full"
)

type PostgresConfig struct {
	ConnString string
}

func NewPostgresConfig(connString string) *PostgresConfig {
	if len(connString) == 0 {
		connString = DefaultPostgresConnString
	}
	cfg := &PostgresConfig{
		ConnString: connString,
	}
	return cfg
}

func (cfg PostgresConfig) Type() Type {
	return Postgres
}

type PostgresBackend struct {
	config *PostgresConfig
	db     *sql.DB // TODO: rename to "db"
	mu     sync.Mutex
}

func NewPostgresBackend(config *PostgresConfig) *PostgresBackend {
	be := &PostgresBackend{
		config: config,
	}
	return be
}

func (be *PostgresBackend) Open() error {
	be.mu.Lock()
	defer be.mu.Unlock()

	if be.db != nil {
		return nil
	}

	db, err := sql.Open("postgres", be.config.ConnString)
	if err != nil {
		return err
	}
	be.db = db

	if err := be.withTx(true, func(pTx *pgTx) error {
		return be.initDB(pTx)
	}); err != nil {
		return err
	}

	return nil
}

func (be *PostgresBackend) Close() error {
	be.mu.Lock()
	defer be.mu.Unlock()

	if be.db == nil {
		return nil
	}

	if err := be.db.Close(); err != nil {
		return err
	}

	be.db = nil

	return nil
}

func (be *PostgresBackend) initDB(pTx *pgTx) error {
	tables := []string{
		TableMetadata,
		TablePackages,
		TableToCrawl,
		TablePendingReferences,
	}
	for _, table := range tables {
		table = be.normalizeTable(table)
		//	id bigserial PRIMARY KEY DEFAULT nextval('serial'),
		_, err := pTx.tx.Exec(fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %s (
	key bytea NOT NULL,
	value bytea NOT NULL
)
`, pq.QuoteIdentifier(table)))
		if err != nil {
			return fmt.Errorf("creating table %q: %s", table, err)
		}
		_, err = pTx.tx.Exec(fmt.Sprintf(
			`CREATE UNIQUE INDEX IF NOT EXISTS %s ON %s (%s)`,
			pq.QuoteIdentifier(fmt.Sprintf("unique_%s_key_index", table)),
			pq.QuoteIdentifier(table),
			pq.QuoteIdentifier("key"),
		))
		if err != nil {
			return fmt.Errorf("creating indexes for table %q: %s", table, err)
		}
	}
	return nil
}

func (be *PostgresBackend) New(name string) (Backend, error) {
	return nil, errors.New("New is not implemented for postgres")
}

func (be *PostgresBackend) Get(table string, key []byte) ([]byte, error) {
	var (
		v []byte
	)
	if err := be.View(func(tx Transaction) error {
		var err error
		if v, err = be.get(tx, table, key); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return v, nil
}

func (be *PostgresBackend) get(tx Transaction, table string, key []byte) ([]byte, error) {
	table = be.normalizeTable(table)
	var (
		pTx = tx.(*pgTx)
		v   []byte
		row = pTx.tx.QueryRow(fmt.Sprintf(`SELECT value FROM %s WHERE key=$1 LIMIT 1`, pq.QuoteIdentifier(table)), key)
	)
	if err := row.Scan(&v); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrKeyNotFound
		}
		return nil, fmt.Errorf("getting key=%q from %v: %s", string(key), table, err)
	}
	//if len(v) == 0 {
	//	return ErrKeyNotFound
	//}
	return v, nil
}

func (be *PostgresBackend) Put(table string, key []byte, value []byte) error {
	return be.withTx(true, func(pTx *pgTx) error {
		return be.put(pTx, table, key, value)
	})
}

func (be *PostgresBackend) put(tx Transaction, table string, key []byte, value []byte) error {
	pTx := tx.(*pgTx)
	_, err := pTx.tx.Exec(
		fmt.Sprintf(
			`
INSERT INTO %s (key, value) VALUES ($1::bytea, $2::bytea)
    ON CONFLICT (key)
    DO UPDATE SET
	value=EXCLUDED.value
`,
			pq.QuoteIdentifier(table),
		),
		key,
		value,
	)
	if err != nil {
		return fmt.Errorf("inserting key=%q into %v: %s", string(key), table, err)
	}
	return nil
}

func (be *PostgresBackend) Delete(table string, keys ...[]byte) error {
	return be.withTx(true, func(pTx *pgTx) error {
		for _, key := range keys {
			_, err := pTx.tx.Exec(fmt.Sprintf(`DELETE FROM %s WHERE key=$1`, pq.QuoteIdentifier(table)), key)
			if err != nil {
				return fmt.Errorf("deleting key=%q from %v: %s", string(key), table, err)
			}
		}
		return nil
	})
}

func (be *PostgresBackend) delete(pTx *pgTx, table string, keys ...[]byte) error {
	return be.withTx(true, func(pTx *pgTx) error {
		return be.delete(pTx, table, keys...)
	})
}

func (be *PostgresBackend) Drop(tables ...string) error {
	return be.withTx(true, func(pTx *pgTx) error {
		for _, table := range tables {
			_, err := pTx.tx.Exec(fmt.Sprintf("DROP TABLE %s", pq.QuoteIdentifier(table)))
			if err != nil {
				return fmt.Errorf("dropping table=%v: %s", table, err)
			}
		}
		if err := be.initDB(pTx); err != nil {
			return err
		}
		return nil
	})
}

func (be *PostgresBackend) Len(table string) (int, error) {
	var n int64
	if err := be.withTx(false, func(pTx *pgTx) error {
		row := pTx.tx.QueryRow(fmt.Sprintf(`SELECT COUNT(*) FROM %s`, pq.QuoteIdentifier(table)))
		if err := row.Scan(&n); err != nil {
			return fmt.Errorf("getting length of table=%v: %s", table, err)
		}
		return nil
	}); err != nil {
		return 0, err
	}
	// TODO : change interface to return int64.
	return int(n), nil
}

func (be *PostgresBackend) Begin(writable bool) (Transaction, error) {
	opts := &sql.TxOptions{}
	if !writable {
		opts.ReadOnly = true
	}
	tx, err := be.db.BeginTx(context.Background(), opts)
	if err != nil {
		return nil, err
	}
	pTx := be.wrapTx(tx)
	return pTx, nil
}

func (be *PostgresBackend) View(fn func(pTx Transaction) error) error {
	return be.withTx(false, func(pTx *pgTx) error {
		return fn(pTx)
	})
}
func (be *PostgresBackend) Update(fn func(pTx Transaction) error) error {
	return be.withTx(true, func(pTx *pgTx) error {
		return fn(pTx)
	})
}

func (be *PostgresBackend) EachRow(table string, fn func(key []byte, value []byte)) error {
	return be.withTx(false, func(pTx *pgTx) error {
		c := pTx.Cursor(table)
		defer c.Close()
		for k, v := c.First().Data(); c.Err() == nil && len(k) > 0; k, v = c.Next().Data() {
			fn(k, v)
		}
		return c.Err()
	})
}

func (be *PostgresBackend) EachRowWithBreak(table string, fn func(key []byte, value []byte) bool) error {
	return be.withTx(false, func(pTx *pgTx) error {
		c := pTx.Cursor(table)
		defer c.Close()
		for k, v := c.First().Data(); c.Err() == nil && len(k) > 0; k, v = c.Next().Data() {
			if !fn(k, v) {
				break
			}
		}
		return c.Err()
	})
}

func (be *PostgresBackend) EachTable(fn func(table string, tx Transaction) error) error {
	return be.View(func(tx Transaction) error {
		var (
			pTx   = tx.(*pgTx)
			table string
		)
		rows, err := pTx.tx.Query(`SELECT table_name FROM information_schema.tables WHERE table_schema='public'`)
		if err != nil {
			return err
		}
		for {
			if err := rows.Scan(&table); err != nil {
				return err
			}
			if err := fn(table, tx); err != nil {
				return err
			}
		}
		return nil
	})
}

func (be *PostgresBackend) withTx(writable bool, fn func(pTx *pgTx) error) error {
	tx, err := be.Begin(true)
	if err != nil {
		return fmt.Errorf("obtaining pTx: %s", err)
	}
	pTx := tx.(*pgTx)
	if err = fn(pTx); err != nil {
		if rbErr := pTx.Rollback(); rbErr != nil {
			return rbErr
		}
		return err
	} else if !writable {
		if err = pTx.Rollback(); err != nil {
			return err
		}
		return nil
	}
	if err = pTx.Commit(); err != nil {
		return err
	}
	return nil

}

func (be *PostgresBackend) normalizeTable(table string) string {
	table = strings.Replace(table, "-", "_", -1)
	return table
}

func (be *PostgresBackend) wrapTx(tx *sql.Tx) *pgTx {
	pTx := &pgTx{
		tx: tx,
		be: be,
	}
	return pTx
}

type pgTx struct {
	tx *sql.Tx
	be *PostgresBackend
}

func (pTx *pgTx) Get(table string, key []byte) ([]byte, error) {
	table = pTx.be.normalizeTable(table)
	return pTx.be.get(pTx, table, key)
}

func (pTx *pgTx) Put(table string, key []byte, value []byte) error {
	table = pTx.be.normalizeTable(table)
	return pTx.be.put(pTx, table, key, value)
}

func (pTx *pgTx) Delete(table string, keys ...[]byte) error {
	table = pTx.be.normalizeTable(table)
	return pTx.be.delete(pTx, table, keys...)
}

func (pTx *pgTx) Commit() error {
	return pTx.tx.Commit()
}

func (pTx *pgTx) Rollback() error {
	return pTx.tx.Rollback()
}

func (pTx *pgTx) Backend() Backend {
	return pTx.be
}

func (pTx *pgTx) Cursor(table string) Cursor {
	table = pTx.be.normalizeTable(table)
	c := newPgCursor(pTx, table)
	return c
}

type pgCursor struct {
	pTx    *pgTx
	table  string
	prefix []byte
	rows   *sql.Rows
	k      []byte
	v      []byte
	err    error
}

func newPgCursor(pTx *pgTx, table string) Cursor {
	c := &pgCursor{
		pTx:   pTx,
		table: table,
	}
	return c
}

func (c *pgCursor) First() Cursor {
	c.setQuery()
	if c.rows.Next() {
		log.Infof("next0")
		c.scan()
	}
	return c
}

func (c *pgCursor) Next() Cursor {
	c.clearRow()
	if c.rows != nil {
		log.Infof("next1")
		if c.rows.Next() {
			c.scan()
		}
	}
	return c
}

func (c *pgCursor) Seek(prefix []byte) Cursor {
	c.prefix = prefix
	c.setQuery()
	c.scan()
	return c
}

func (c *pgCursor) Data() (key []byte, value []byte) {
	return c.k, c.v
}

func (c *pgCursor) Close() {
	c.discard()
}

func (c *pgCursor) setQuery() {
	c.discard()
	var err error
	if len(c.prefix) > 0 {
		c.rows, err = c.pTx.tx.Query(
			fmt.Sprintf(`SELECT "key", "value" FROM %s WHERE "key" LIKE $1 ORDER BY "key" ASC`, pq.QuoteIdentifier(c.table)),
			[]byte(fmt.Sprintf("%v%%", string(c.prefix))),
		)
	} else {
		c.rows, err = c.pTx.tx.Query(
			fmt.Sprintf(`SELECT "key", "value" FROM %s ORDER BY "key" ASC`, pq.QuoteIdentifier(c.table)),
		)
	}
	if err != nil {
		log.Errorf("PG seek error: %s (continuing on..)", err)
		c.err = err
		return
	}
}

// discard clears out current rows, if any.
func (c *pgCursor) discard() {
	if c.rows != nil {
		c.rows.Close()
		c.rows = nil

		c.clearRow()

		// Clear out potential error state.
		c.err = nil
	}
}

func (c *pgCursor) clearRow() {
	c.k = nil
	c.v = nil
}

func (c *pgCursor) scan() {
	if err := c.rows.Scan(&c.k, &c.v); err != nil {
		log.Errorf("PG key scan error: %s (continuing on..)", err)
		c.err = err
		return
	}
}

func (c *pgCursor) Err() error {
	return c.err
}
