package collector

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/prometheus/client_golang/prometheus"

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

// minimum PostgreSQL version
func (ScrapeBgWriter) Version() int {
	return 0
}

// scrape type
func (ScrapeBgWriter) Type() ScrapeType {
	return SCRAPEGLOBAL
}

// Scrape collects data from database connection and sends it over channel as prometheus metric.
func (ScrapeBgWriter) Scrape(ctx context.Context, db *pg.DB, ch chan<- prometheus.Metric) error {
	statBgwriter := &models.PgStatBgWriter{}
	if err := db.ModelContext(ctx, statBgwriter).Select(); err != nil {
		return err
	}

	return statBgwriter.ToMetrics(namespace, bgwriter, ch)
}
