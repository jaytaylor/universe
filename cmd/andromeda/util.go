package main

import (
	"fmt"
	"strings"

	"github.com/gogo/protobuf/proto"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"jaytaylor.com/andromeda/crawler"
	"jaytaylor.com/andromeda/db"
	"jaytaylor.com/andromeda/domain"
)

func newUtilCmd() *cobra.Command {
	utilCmd := &cobra.Command{
		Use:     "util",
		Aliases: []string{"utility", "utils", "utilities"},
		Short:   "Utility tools and functions",
		Long:    "Utility tools and functions",
	}

	utilCmd.AddCommand(
		newRepoRootCmd(),
		newRebuildDBCmd(),
	)

	return utilCmd
}

// newRepoRootCmd TODO: move this to util sub-command.
func newRepoRootCmd() *cobra.Command {
	repoRootCmd := &cobra.Command{
		Use:     "repo-root <package-path>",
		Aliases: []string{"reporoot", "rr"},
		Short:   "Package repository root lookup",
		Long:    "Administrative utilithy to lookup the repository root for a package",
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			rr, err := crawler.PackagePathToRepoRoot(args[0])
			if err != nil {
				log.Fatalf("%s", err)
			}
			if err := emitJSON(rr); err != nil {
				log.Fatalf("main: %s", err)
			}
		},
	}
	return repoRootCmd
}

func newRebuildDBCmd() *cobra.Command {
	filters := map[string]db.KeyValueFilterFunc{
		"normalizeSubPackageKeys": func(table []byte, k []byte, v []byte) ([]byte, []byte) {
			if string(table) != db.TablePackages {
				return k, v
			}
			pkg := &domain.Package{}
			if err := proto.Unmarshal(v, pkg); err != nil {
				log.Fatalf("Unexpected problem unmarshaling protobuf for key=%v: %s", string(k), err)
			}

			for subPkgPath, _ := range pkg.Data.SubPackages {
				if strings.Contains(subPkgPath, "/vendor/") || strings.Contains(subPkgPath, "Godep/_workspace/") {
					delete(pkg.Data.SubPackages, subPkgPath)
				}
			}
			pkg.NormalizeSubPackageKeys()

			v, err := proto.Marshal(pkg)
			if err != nil {
				log.Fatalf("Unexpected problem marshaling protobuf for key=%v: %s", string(k), err)
			}
			return k, v
		},
		"clearHistories": func(table []byte, k []byte, v []byte) ([]byte, []byte) {
			if string(table) != db.TablePackages {
				return k, v
			}
			pkg := &domain.Package{}
			if err := proto.Unmarshal(v, pkg); err != nil {
				log.Fatalf("Unexpected problem unmarshaling protobuf for key=%v: %s", string(k), err)
			}
			pkg.History = []*domain.PackageCrawl{pkg.History[0]}

			v, err := proto.Marshal(pkg)
			if err != nil {
				log.Fatalf("Unexpected problem marshaling protobuf for key=%v: %s", string(k), err)
			}
			return k, v
		},
	}

	filterNames := []string{}
	for filterName, _ := range filters {
		filterNames = append(filterNames, filterName)
	}

	var (
		skipKV bool
		skipQ  bool
	)

	rebuildDBCmd := &cobra.Command{
		Use:     "rebuild-db",
		Aliases: []string{"rebuild"},
		Short:   "Rebuilds the database",
		Long:    "Rebuilds the entire database",
		PreRun: func(_ *cobra.Command, _ []string) {
			if len(RebuildDBDriver) == 0 {
				log.Fatal("rebuild-db-driver value is required")
			}
			if len(RebuildDBFile) == 0 {
				log.Fatal("rebuild-db value is required")
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			dbCfg := db.NewConfig(DBDriver, DBFile)
			if err := db.WithClient(dbCfg, func(dbClient db.Client) error {
				newCfg := db.NewConfig(RebuildDBDriver, RebuildDBFile)

				applyFilters := []db.KeyValueFilterFunc{}
				for _, name := range RebuildDBFilters {
					if filter, ok := filters[name]; ok {
						applyFilters = append(applyFilters, filter)
					} else {
						log.Fatalf("Unrecognized filter %q", name)
					}
				}

				if skipKV {
					applyFilters = append(applyFilters, db.SkipKVFilter)
				}
				if skipQ {
					applyFilters = append(applyFilters, db.SkipQFilter)
				}

				return db.WithClient(newCfg, func(newClient db.Client) error {
					return dbClient.RebuildTo(newClient, applyFilters...)
				})
			}); err != nil {
				log.Fatal(err)
			}
		},
	}

	rebuildDBCmd.Flags().StringVarP(&RebuildDBDriver, "rebuild-driver", "", RebuildDBDriver, "Target destination DB driver")
	rebuildDBCmd.Flags().StringVarP(&RebuildDBFile, "rebuild-db", "", RebuildDBFile, "Target destination filename or db connection string")
	rebuildDBCmd.Flags().StringSliceVarP(&RebuildDBFilters, "rebuild-filters", "", RebuildDBFilters, fmt.Sprintf("Comma-delimited list of filter function(s) to apply; available filters: %s", strings.Join(filterNames, ", ")))
	rebuildDBCmd.Flags().IntVarP(&db.RebuildBatchSize, "batch-size", "", db.RebuildBatchSize, "Transaction batch size")
	rebuildDBCmd.Flags().BoolVar(&skipKV, "skip-kv", skipKV, "Skip Key/Value stores")
	rebuildDBCmd.Flags().BoolVar(&skipQ, "skip-q", skipQ, "Skip Queue stores")

	return rebuildDBCmd
}
