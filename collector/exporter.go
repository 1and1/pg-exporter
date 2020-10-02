package collector

import (
	"context"
	"sync"
	"time"

	"github.com/go-pg/pg/v9"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
)

// Metric name parts.
const (
	// Subsystem(s).
	exporter = "exporter"
)

// SQL queries and parameters.
const ()

// calculated vars
var (
	pgversion int
)

// Metric descriptors.
var (
	scrapeDurationDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, exporter, "collector_duration_seconds"),
		"Collector time duration.",
		[]string{"collector"}, nil,
	)
)

// Verify if Exporter implements prometheus.Collector
var _ prometheus.Collector = (*Exporter)(nil)

// Exporter collects MySQL metrics. It implements prometheus.Collector.
type Exporter struct {
	ctx       context.Context
	dsn       string
	pgoptions *pg.Options
	scrapers  []Scraper
	metrics   Metrics
}

// New returns a PostgreSQL exporter for the provided DSN.
func New(ctx context.Context, pgoptions *pg.Options, metrics Metrics, scrapers []Scraper) *Exporter {
	// Database exporters should only open one connection
	pgoptions.PoolSize = 1
	// set the lock timeout, to minimice possible production impact

	// TODO: this triggers a segfault in some cases, we need to investigate why later
	/*
	   oldOnconnect := pgoptions.OnConnect
	   newOnconnect := func(conn *pg.Conn) error {
	       conn.ExecContext(ctx, "SET lock_timeout = 5")
	       if oldOnconnect != nil {
	           return oldOnconnect(conn)
	       }
	       return nil
	   }

	   pgoptions.OnConnect = newOnconnect
	*/
	return &Exporter{
		ctx:       ctx,
		pgoptions: pgoptions,
		scrapers:  scrapers,
		metrics:   metrics,
	}
}

// Describe implements prometheus.Collector.
func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- e.metrics.TotalScrapes.Desc()
	ch <- e.metrics.Error.Desc()
	e.metrics.ScrapeErrors.Describe(ch)
	ch <- e.metrics.PgSQLUp.Desc()
}

// Collect implements prometheus.Collector.
func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	e.scrape(e.ctx, ch)

	ch <- e.metrics.TotalScrapes
	ch <- e.metrics.Error
	e.metrics.ScrapeErrors.Collect(ch)
	ch <- e.metrics.PgSQLUp
}

func (e *Exporter) scrape(ctx context.Context, ch chan<- prometheus.Metric) {
	e.metrics.TotalScrapes.Inc()

	scrapeTime := time.Now()

	db := pg.Connect(e.pgoptions)
	defer db.Close()

	// get the database version
	if _, err := db.QueryContext(ctx, pg.Scan(&pgversion), "SHOW server_version_num"); err != nil {
		log.Errorln("Error getting PostgreSQL version:", err)
		e.metrics.PgSQLUp.Set(0)
		e.metrics.Error.Set(1)
		return
	}

	// update our requested db list
	if err := updateDatabaseList(ctx, db); err != nil {
		log.Errorln("error updating database list:", err)
		e.metrics.PgSQLUp.Set(0)
		e.metrics.Error.Set(1)
		return
	}

	e.metrics.PgSQLUp.Set(1)
	e.metrics.Error.Set(0)

	ch <- prometheus.MustNewConstMetric(scrapeDurationDesc, prometheus.GaugeValue, time.Since(scrapeTime).Seconds(), "connection")

	var wg sync.WaitGroup

	// collect the statistics with a global view
	for _, scraper := range e.scrapers {
		if scraper.Type() != SCRAPEGLOBAL {
			continue
		}
		if pgversion < scraper.Version() {
			continue
		}

		wg.Add(1)
		go func(scraper Scraper) {
			defer wg.Done()
			label := "collect." + scraper.Name()
			scrapeTime := time.Now()
			if err := scraper.Scrape(ctx, db, ch); err != nil {
				log.Errorln("Error scraping for "+label+":", err)
				e.metrics.ScrapeErrors.WithLabelValues(label).Inc()
				e.metrics.Error.Set(1)
			}
			ch <- prometheus.MustNewConstMetric(scrapeDurationDesc, prometheus.GaugeValue, time.Since(scrapeTime).Seconds(), label)
		}(scraper)
	}

	// collect the database scoped statistics
	for _, dbname := range collectDatabases {
		// copy the contents of pgoptions
		// we don't require a deep copy, as we only change a string value at the top level
		dboptions := *e.pgoptions
		dboptions.Database = dbname
		// TODO: the exporter should open only one connection to the database instance,
		// as we need one connection per database, this can increase here to n + 1 where n is the number of databases
		// to scrape
		localdb := pg.Connect(&dboptions)
		defer localdb.Close()
		for _, scraper := range e.scrapers {
			if scraper.Type() != SCRAPELOCAL {
				continue
			}
			if pgversion < scraper.Version() {
				continue
			}

			wg.Add(1)
			go func(scraper Scraper, dbname string, db *pg.DB) {
				defer wg.Done()
				label := "collect." + scraper.Name() + "." + dbname
				scrapeTime := time.Now()
				if err := scraper.Scrape(ctx, localdb, ch); err != nil {
					log.Errorf("Error scraping for %s: %v", label, err)
					e.metrics.ScrapeErrors.WithLabelValues(label).Inc()
					e.metrics.Error.Set(1)
				}
				ch <- prometheus.MustNewConstMetric(scrapeDurationDesc, prometheus.GaugeValue,
					time.Since(scrapeTime).Seconds(), label)
			}(scraper, dbname, db)
		}
	}
	wg.Wait()
}

// Metrics represents exporter metrics which values can be carried between http requests.
type Metrics struct {
	TotalScrapes prometheus.Counter
	ScrapeErrors *prometheus.CounterVec
	Error        prometheus.Gauge
	PgSQLUp      prometheus.Gauge
}

// NewMetrics creates new Metrics instance.
func NewMetrics() Metrics {
	subsystem := exporter
	return Metrics{
		TotalScrapes: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "scrapes_total",
			Help:      "Total number of times PostgreSQL was scraped for metrics.",
		}),
		ScrapeErrors: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "scrape_errors_total",
			Help:      "Total number of times an error occurred scraping a PostgreSQL.",
		}, []string{"collector"}),
		Error: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "last_scrape_error",
			Help:      "Whether the last scrape of metrics from PostgreSQL resulted in an error (1 for error, 0 for success).",
		}),
		PgSQLUp: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "up",
			Help:      "Whether the PostgreSQL server is up.",
		}),
	}
}
