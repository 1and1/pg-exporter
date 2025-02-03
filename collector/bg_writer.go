package collector

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/uptrace/bun"

	"github.com/1and1/pg-exporter/collector/models"
)

const (
	// subsystem
	bgwriter = "bgwriter"
)

// ScrapeBgWriter scrapes from pg_stat_bgwriter
type ScrapeBgWriter struct{}

// Name of the Scraper
func (ScrapeBgWriter) Name() string {
	return "pg_stat_bgwriter"
}

// Help describes the role of the Scraper
func (ScrapeBgWriter) Help() string {
	return "Collect from pg_stat_bgwriter"
}

// Version returns minimum PostgreSQL version
func (ScrapeBgWriter) Version() int {
	return 0
}

// Type returns the scrape type
func (ScrapeBgWriter) Type() ScrapeType {
	return SCRAPEGLOBAL
}

// Scrape collects data from database connection and sends it over channel as prometheus metric.
func (ScrapeBgWriter) Scrape(ctx context.Context, db *bun.DB, ch chan<- prometheus.Metric) error {
    if pgversion < 170000 {
        statBgwriter := &models.PgStatBgWriter{}

        if err := db.NewSelect().Model(statBgwriter).Scan(ctx); err != nil {
            return err
        }

        return statBgwriter.ToMetrics(namespace, bgwriter, ch)
    } else {

	    statBgwriter := &models.PgStatBgWriter17{}
        if err := db.NewSelect().Model(statBgwriter).Scan(ctx); err != nil {
            return err
        }

        return statBgwriter.ToMetrics(namespace, bgwriter, ch)
    }
}
