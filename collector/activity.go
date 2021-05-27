package collector

import (
	"context"
	"fmt"

	"github.com/go-pg/pg/v9"
	"github.com/prometheus/client_golang/prometheus"
	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/1and1/pg-exporter/collector/models"
)

const (
	// subsystem
	activity = "sessions"
)

// additional command line options
var (
	withUsername = kingpin.Flag("collect.pg_stat_activity.with_username",
		"Include username in session statistics").
		Default("false").Bool()
	withApplicationName = kingpin.Flag("collect.pg_stat_activity.with_appname",
		"Include application name in session statistics").
		Default("false").Bool()
	withClientAddr = kingpin.Flag("collect.pg_stat_activity.with_clientaddr",
		"Include application name in session statistics").
		Default("false").Bool()
	withState = kingpin.Flag("collect.pg_stat_activity.with_state",
		"Include session state in session statistics").
		Default("true").Bool()
	withWaitEventType = kingpin.Flag("collect.pg_stat_activity.with_wait_type",
		"Include wait_event_type in session statistics").
		Default("false").Bool()
	withBackendType = kingpin.Flag("collect.pg_stat_activity.with_backend_type",
		"Include backend_type in session statistics").
		Default("false").Bool()
)

// ScrapeActivity scrapes from pg_stat_bgwriter
type ScrapeActivity struct{}

// Name of the Scraper
func (ScrapeActivity) Name() string {
	return "pg_stat_activity"
}

// Help describes the role of the Scraper
func (ScrapeActivity) Help() string {
	return "Collect from pg_stat_activity"
}

// Version returns minimum PostgreSQL version
func (ScrapeActivity) Version() int {
	return 0
}

// Type returns the scrape type
func (ScrapeActivity) Type() ScrapeType {
	return SCRAPEGLOBAL
}

// Scrape collects data from database connection and sends it over channel as prometheus metric.
func (ScrapeActivity) Scrape(ctx context.Context, db *pg.DB, ch chan<- prometheus.Metric) error {
	// we create a query based on the given commandline flags
	columns := ""
	if *withUsername {
		columns += ", usename"
	}

	if *withApplicationName {
		columns += ", application_name"
	}

	if *withClientAddr {
		columns += ", client_addr"
	}

	if *withState {
		columns += ", state"
	}

	if *withWaitEventType {
		columns += ", wait_event_type"
	}

	if pgversion >= 110000 && *withBackendType {
		columns += ", backend_type"
	}

	qs := fmt.Sprintf(`SELECT datid, datname %s, count(1) as connections FROM`+
		` pg_stat_activity WHERE datname IN (?) AND state IS NOT NULL and state != '' GROUP BY datid, datname %s`,
		columns, columns)

	var statActivity models.PgStatActivitySlice
	if _, err := db.QueryContext(ctx, &statActivity, qs, pg.In(collectDatabases)); err != nil {
		return err
	}
	return statActivity.ToMetrics(namespace, activity, ch)
}
