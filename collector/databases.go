package collector

import (
	"context"

	"github.com/uptrace/bun"
	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/1and1/pg-exporter/collector/models"
)

// commandline flags
var (
	includeDatabases = kingpin.Flag(
		"collect.include.database",
		"Database to include in the collection. If defined only the given databases are scraped. "+
			"Can be defined multiple times",
	).Default().Strings()

	excludeDatabases = kingpin.Flag(
		"collect.exclude.database",
		"Database to exclude from the collection. Only used if collect.include.database is not given. "+
			"Can be defined multiple times",
	).Default("template0", "template1", "postgres").Strings()
)

// calculated vars
var (
	collectDatabases []string
)

func updateDatabaseList(ctx context.Context, db *bun.DB) error {
	// prepare lookup maps
	prepareLookup := func(in *[]string) map[string]bool {
		lookup := make(map[string]bool)
		for _, name := range *in {
			lookup[name] = true
		}
		return lookup
	}
	includeLookup := prepareLookup(includeDatabases)
	excludeLookup := prepareLookup(excludeDatabases)

	var dblist []string
	var databases []models.PgDatabase
	if err := db.NewSelect().Model(&databases).Scan(ctx); err != nil {
		return err
	}
	// iterate over all databases and check if we include or exclude them
	for _, pgdb := range databases {
		if len(includeLookup) > 0 {
			if includeLookup[pgdb.Datname] {
				dblist = append(dblist, pgdb.Datname)
			}
		} else {
			if !excludeLookup[pgdb.Datname] {
				dblist = append(dblist, pgdb.Datname)
			}
		}
	}

	collectDatabases = dblist

	return nil
}
