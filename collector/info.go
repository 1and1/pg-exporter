package collector

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	// subsystem
	info = "info"
)

// ScrapeSettings scrapes from pg_settings
type ScrapeInfo struct{}

// Name of the Scraper
func (ScrapeInfo) Name() string {
	return "info"
}

// Help describes the role of the Scraper
func (ScrapeInfo) Help() string {
	return "Collect postgresql information"
}

// minimum PostgreSQL version
func (ScrapeInfo) Version() int {
	return 0
}

// scrape type
func (ScrapeInfo) Type() ScrapeType {
	return SCRAPEGLOBAL
}

// Scrape collects data from database connection and sends it over channel as prometheus metric.
func (ScrapeInfo) Scrape(ctx context.Context, db *pg.DB, ch chan<- prometheus.Metric) error {
	// we reuse the definition from pg_settings
	var settingsRes []pgSetting
	if err := db.ModelContext(ctx, &settingsRes).WhereIn("name IN (?)", []string{"server_version"}).Select(); err != nil {
		return err
	}
	labels := make(map[string]string)
	for _, setting := range settingsRes {
		labels[setting.Name] = setting.Setting
	}
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", info), "information about the PostgreSQL instance", nil, labels,
		), prometheus.GaugeValue, 1,
	)
	return nil
}
