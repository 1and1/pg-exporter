package collector

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/uptrace/bun"

	"github.com/1and1/pg-exporter/collector/models"
)

const (
	// subsystem
	checkpointer = "checkpointer"
)

// ScrapeCheckpointer scrapes from pg_stat_checkpointer
type ScrapeCheckpointer struct{}

// Name of the Scraper
func (ScrapeCheckpointer) Name() string {
	return "pg_stat_checkpointer"
}

// Help describes the role of the Scraper
func (ScrapeCheckpointer) Help() string {
	return "Collect from pg_stat_checkpointer"
}

// Version returns minimum PostgreSQL version
func (ScrapeCheckpointer) Version() int {
	return 0
}

// Type returns the scrape type
func (ScrapeCheckpointer) Type() ScrapeType {
	return SCRAPEGLOBAL
}

// Scrape collects data from database connection and sends it over channel as prometheus metric.
func (ScrapeCheckpointer) Scrape(ctx context.Context, db *bun.DB, ch chan<- prometheus.Metric) error {
    if pgversion >= 170000 {
	    statCheckpointer := &models.PgStatCheckpointer{}
        if err := db.NewSelect().Model(statCheckpointer).Scan(ctx); err != nil {
            return err
        }

        return statCheckpointer.ToMetrics(namespace, checkpointer, ch)
    } else {
        return nil
    }
}
