package collector

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/prometheus/common/log"
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

func updateDatabaseList(ctx context.Context, db *pg.DB) error {
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
	// iterate over all databases and check if we include or exclude them
	err := db.ModelContext(ctx, (*models.PgDatabase)(nil)).ForEach(func(pgdb *models.PgDatabase) error {
		if len(includeLookup) > 0 {
			if includeLookup[pgdb.Datname] {
				dblist = append(dblist, pgdb.Datname)
			}
		} else {
			if !excludeLookup[pgdb.Datname] {
				dblist = append(dblist, pgdb.Datname)
			}
		}
		return nil
	})

	if err != nil {
		return err
	}

	collectDatabases = dblist

	log.Debugf("effective database list: %v", collectDatabases)

	return nil
}
