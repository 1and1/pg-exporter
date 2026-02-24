package collector

import (
	"context"

    "fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/uptrace/bun"

	"github.com/1and1/pg-exporter/collector/models"
)

const (
	// subsystem
	io = "io"
)

// ScrapeWal scrapes from pg_stat_io
type ScrapeIo struct{}

// Name of the Scraper
func (ScrapeIo) Name() string {
	return "pg_stat_io"
}

// Help describes the role of the Scraper
func (ScrapeIo) Help() string {
	return "Collect from pg_stat_io"
}

// Version returns minimum PostgreSQL version
func (ScrapeIo) Version() int {
	return 160000
}

// Type returns the scrape type
func (ScrapeIo) Type() ScrapeType {
	return SCRAPEGLOBAL
}

// Scrape collects data from database connection and sends it over channel as prometheus metric.
func (ScrapeIo) Scrape(ctx context.Context, db *bun.DB, ch chan<- prometheus.Metric) error {
    // var qs string
    columns := ""
    if pgversion < 180000 {
            columns += "backend_type, context, evictions, extend_time, extends, fsync_time, fsyncs, hits, object, op_bytes, read_time, reads, reuses, stats_reset, write_time, writeback_time, writebacks, writes"    }

    if pgversion >= 180000 {
            columns += "backend_type, context, evictions, extend_bytes, extend_time, extends, fsync_time, fsyncs, hits, object, read_bytes, read_time, reads, reuses, stats_reset, write_bytes, write_time, writeback_time, writebacks, writes"
    }

    qs := fmt.Sprintf(`SELECT %s FROM pg_stat_io`, columns)

    var statIo models.PgStatIo
    rows, err := db.QueryContext(ctx, qs)
    if err != nil {
        return err
    }
    if err := db.ScanRows(ctx, rows, &statIo); err != nil {
        return err
    }
    return statIo.ToMetrics(namespace, io, ch)
}
