package collector

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/uptrace/bun"
)

const (
	// subsystem
	info = "info"
)

// ScrapeInfo scrapes generic PostgreSQL information
type ScrapeInfo struct{}

// Name of the Scraper
func (ScrapeInfo) Name() string {
	return "info"
}

// Help describes the role of the Scraper
func (ScrapeInfo) Help() string {
	return "Collect postgresql information"
}

// Version returns minimum PostgreSQL version
func (ScrapeInfo) Version() int {
	return 0
}

// Type returns the scrape type
func (ScrapeInfo) Type() ScrapeType {
	return SCRAPEGLOBAL
}

// Scrape collects data from database connection and sends it over channel as prometheus metric.
func (ScrapeInfo) Scrape(ctx context.Context, db *bun.DB, ch chan<- prometheus.Metric) error {
	// we reuse the definition from pg_settings
	var settingsRes []pgSetting
	if err := db.NewSelect().Model(&settingsRes).Where("name IN (?)", bun.In([]string{"server_version"})).Scan(ctx); err != nil {
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
