package collector

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/uptrace/bun"

	"github.com/1and1/pg-exporter/collector/models"
)

const (
	// subsystem
	replication_slots = "replication_slots"
)

// ScrapeReplicationSlots scrapes from pg_stat_replication_slots
type ScrapeReplicationSlots struct{}

// Name of the Scraper
func (ScrapeReplicationSlots) Name() string {
	return "pg_stat_replication_slots"
}

// Help describes the role of the Scraper
func (ScrapeReplicationSlots) Help() string {
	return "Collect from pg_stat_replication_slots"
}

// Version returns minimum PostgreSQL version
func (ScrapeReplicationSlots) Version() int {
	return 140000
}

// Type returns the scrape type
func (ScrapeReplicationSlots) Type() ScrapeType {
	return SCRAPEGLOBAL
}

// Scrape collects data from database connection and sends it over channel as prometheus metric.
func (ScrapeReplicationSlots) Scrape(ctx context.Context, db *bun.DB, ch chan<- prometheus.Metric) error {
	statReplicationSlots := &models.PgStatReplicationSlotsSlice{}
	if err := db.NewSelect().Model(statReplicationSlots).Scan(ctx); err != nil {
		return err
	}

	return statReplicationSlots.ToMetrics(namespace, replication_slots, ch)
}
