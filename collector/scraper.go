package collector

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/prometheus/client_golang/prometheus"
)

type ScrapeType int

const (
	SCRAPEGLOBAL ScrapeType = iota
	SCRAPELOCAL
)

// Scraper is minimal interface that let's you add new prometheus metrics to pg_exporter.
type Scraper interface {
	// Name of the Scraper. Should be unique.
	Name() string

	// Help describes the role of the Scraper.
	Help() string

	// Version of PostgreSQL from which scraper is available.
	Version() int

	// Type defines the scrape type
	Type() ScrapeType

	// Scrape collects data from database connection and sends it over channel as prometheus metric.
	Scrape(ctx context.Context, db *pg.DB, ch chan<- prometheus.Metric) error
}
