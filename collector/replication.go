package collector

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/1and1/pg-exporter/collector/models"
)

const (
	// subsystem
	replication = "replication"
)

// ScrapeReplication scrapes from pg_stat_replication
type ScrapeReplication struct{}

// Name of the Scraper
func (ScrapeReplication) Name() string {
	return "pg_stat_replication"
}

// Help describes the role of the Scraper
func (ScrapeReplication) Help() string {
	return "Collect from pg_stat_replication"
}

// Version returns minimum PostgreSQL version
func (ScrapeReplication) Version() int {
	return 90200
}

// Type returns the scrape type
func (ScrapeReplication) Type() ScrapeType {
	return SCRAPEGLOBAL
}

// Scrape collects data from database connection and sends it over channel as prometheus metric.
func (ScrapeReplication) Scrape(ctx context.Context, db *pg.DB, ch chan<- prometheus.Metric) error {
	var qs string
	if pgversion < 100000 {
		qs = `SELECT pid,
                   application_name,
                   client_addr,
                   backend_xmin,
                   pg_xlog_location_diff(
                           CASE WHEN pg_is_in_recovery() THEN pg_last_xlog_receive_location() ELSE pg_current_xlog_location() END,
                           sent_location)   AS sent_lag_bytes,
                   pg_xlog_location_diff(
                           CASE WHEN pg_is_in_recovery() THEN pg_last_xlog_receive_location() ELSE pg_current_xlog_location() END,
                           write_location)   AS write_lag_bytes,
                   pg_xlog_location_diff(
                           CASE WHEN pg_is_in_recovery() THEN pg_last_xlog_receive_location() ELSE pg_current_xlog_location() END,
                           flush_location)  AS flush_lag_bytes,
                   pg_xlog_location_diff(
                           CASE WHEN pg_is_in_recovery() THEN pg_last_xlog_receive_location() ELSE pg_current_xlog_location() END,
                           replay_location) AS replay_lag_bytes
            FROM pg_stat_replication;`
	} else {
		qs = `SELECT pid,
                   application_name,
                   client_addr,
                   backend_xmin,
                   pg_wal_lsn_diff(CASE WHEN pg_is_in_recovery() THEN pg_last_wal_receive_lsn() ELSE pg_current_wal_lsn() END,
                                   sent_lsn)      AS sent_lag_bytes,
                   pg_wal_lsn_diff(CASE WHEN pg_is_in_recovery() THEN pg_last_wal_receive_lsn() ELSE pg_current_wal_lsn() END,
                                   flush_lsn)     AS flush_lag_bytes,
                   pg_wal_lsn_diff(CASE WHEN pg_is_in_recovery() THEN pg_last_wal_receive_lsn() ELSE pg_current_wal_lsn() END,
                                   write_lsn)   AS write_lag_bytes,
                   pg_wal_lsn_diff(CASE WHEN pg_is_in_recovery() THEN pg_last_wal_receive_lsn() ELSE pg_current_wal_lsn() END,
                                   replay_lsn)    AS replay_lag_bytes,
                   extract(EPOCH FROM write_lag)  AS write_lag,
                   extract(EPOCH FROM flush_lag)  AS flush_lag,
                   extract(EPOCH FROM replay_lag) AS replay_lag
            FROM pg_stat_replication;`
	}
	var statReplication models.PgStatReplicationSlice
	if _, err := db.QueryContext(ctx, &statReplication, qs); err != nil {
		return err
	}
	return statReplication.ToMetrics(namespace, replication, ch)
}
