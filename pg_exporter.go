package main

//go:generate go run github.com/1and1/pg-exporter/gen -i ./collector/models -o collector/models

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
	"github.com/prometheus/common/version"
	"github.com/uptrace/bun/driver/pgdriver"
	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/1and1/pg-exporter/collector"
)

var (
	listenAddress = kingpin.Flag(
		"web.listen-address",
		"Address to listen on for web interface and telemetry.",
	).Default(":9135").String()
	metricPath = kingpin.Flag(
		"web.telemetry-path",
		"Path under which to expose metrics.",
	).Default("/metrics").String()
	timeoutOffset = kingpin.Flag(
		"timeout-offset",
		"Offset to subtract from timeout in seconds.",
	).Default("0.25").Float64()
	pgoptions []pgdriver.Option
)

var scrapers = map[collector.Scraper]bool{
	collector.ScrapeInfo{}:              true,
	collector.ScrapeBgWriter{}:          true,
	collector.ScrapeDatabase{}:          true,
	collector.ScrapeDatabaseConflicts{}: true,
	collector.ScrapeActivity{}:          true,
	collector.ScrapeTables{}:            true,
	collector.ScrapeSettings{}:          true,
	collector.ScrapeLocks{}:             true,
	collector.ScrapeArchiver{}:          true,
	collector.ScrapeReplication{}:       true,
	collector.ScrapeReplicationSlots{}:  true,
	collector.ScrapePreparedXacts{}:     true,
	collector.ScrapeIOTables{}:          true,
	collector.ScrapeWal{}:               true,
	collector.ScrapeFrozenXid{}:         true,
	collector.ScrapeStatements{}:        false,
}

func init() {
	prometheus.MustRegister(version.NewCollector("pg_exporter"))
}

func newHandler(metrics collector.Metrics, scrapers []collector.Scraper) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		filteredScrapers := scrapers
		params := r.URL.Query()["collect[]"]
		// Use request context for cancellation when connection gets closed.
		ctx := r.Context()
		// If a timeout is configured via the Prometheus header, add it to the context.
		if v := r.Header.Get("X-Prometheus-Scrape-Timeout-Seconds"); v != "" {
			timeoutSeconds, err := strconv.ParseFloat(v, 64)
			if err != nil {
				log.Errorf("Failed to parse timeout from Prometheus header: %s", err)
			} else {
				if *timeoutOffset >= timeoutSeconds {
					// Ignore timeout offset if it doesn't leave time to scrape.
					log.Errorf(
						"Timeout offset (--timeout-offset=%.2f) should be lower than prometheus scrape time (X-Prometheus-Scrape-Timeout-Seconds=%.2f).",
						*timeoutOffset,
						timeoutSeconds,
					)
				} else {
					// Subtract timeout offset from timeout.
					timeoutSeconds -= *timeoutOffset
				}
				// Create new timeout context with request context as parent.
				var cancel context.CancelFunc
				ctx, cancel = context.WithTimeout(ctx, time.Duration(timeoutSeconds*float64(time.Second)))
				defer cancel()
				// Overwrite request with timeout context.
				r = r.WithContext(ctx)
			}
		}
		log.Debugln("collect query:", params)

		// Check if we have some "collect[]" query parameters.
		if len(params) > 0 {
			filters := make(map[string]bool)
			for _, param := range params {
				filters[param] = true
			}

			filteredScrapers = nil
			for _, scraper := range scrapers {
				if filters[scraper.Name()] {
					filteredScrapers = append(filteredScrapers, scraper)
				}
			}
		}

		registry := prometheus.NewRegistry()
		registry.MustRegister(collector.New(ctx, pgoptions, metrics, filteredScrapers))

		gatherers := prometheus.Gatherers{
			prometheus.DefaultGatherer,
			registry,
		}
		// Delegate http serving to Prometheus client library, which will call collector.Collect.
		h := promhttp.HandlerFor(gatherers, promhttp.HandlerOpts{})
		h.ServeHTTP(w, r)

	}
}

func main() {
	// Generate ON/OFF flags for all scrapers.
	scraperFlags := map[collector.Scraper]*bool{}
	for scraper, enabledByDefault := range scrapers {
		defaultOn := "false"
		if enabledByDefault {
			defaultOn = "true"
		}

		f := kingpin.Flag(
			"collect."+scraper.Name(),
			scraper.Help(),
		).Default(defaultOn).Bool()

		scraperFlags[scraper] = f
	}

	// Parse flags.
	log.AddFlags(kingpin.CommandLine)
	kingpin.Version(version.Print("pg_exporter"))
	kingpin.HelpFlag.Short('h')
	kingpin.Parse()

	// landingPage contains the HTML served at '/'.
	// TODO: Make this nicer and more informative.
	var landingPage = []byte(`<html>
<head><title>PostgreSQL exporter</title></head>
<body>
<h1>PostgreSQL exporter</h1>
<p><a href='` + *metricPath + `'>Metrics</a></p>
</body>
</html>
`)

	log.Infoln("Starting pg_exporter", version.Info())
	log.Infoln("Build context", version.BuildContext())

	// if we have a dsn, we use this for the connection
	if dsn := os.Getenv("DATA_SOURCE_NAME"); dsn != "" {
		pgoptions = append(pgoptions,
			pgdriver.WithDSN(dsn), // parse the DSN
		)
	} else if pghost := os.Getenv("PGHOST"); pghost != "" {
		// if we have a pghost, we check if it starts with /, if so, set to unix mode
		if strings.HasPrefix(pghost, "/") {
			port := os.Getenv("PGPORT")
			if port == "" {
				port = "5432"
			}
			if matched, _ := regexp.MatchString(`.*?/\.s\.PGSQL\.\d+`, pghost); !matched {
				pghost = fmt.Sprintf(`%s/.s.PGSQL.%s`, pghost, port)
			}
			pgoptions = append(pgoptions,
				pgdriver.WithNetwork("unix"),
				pgdriver.WithAddr(pghost),
				pgdriver.WithInsecure(true),
			)
		}
	}
	// Register only scrapers enabled by flag.
	log.Infof("Enabled scrapers:")
	enabledScrapers := []collector.Scraper{}
	for scraper, enabled := range scraperFlags {
		if *enabled {
			log.Infof(" --collect.%s", scraper.Name())
			enabledScrapers = append(enabledScrapers, scraper)
		}
	}
	handlerFunc := newHandler(collector.NewMetrics(), enabledScrapers)
	http.Handle(*metricPath, promhttp.InstrumentMetricHandler(prometheus.DefaultRegisterer, handlerFunc))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write(landingPage)
	})

	log.Infoln("Listening on", *listenAddress)
	log.Fatal(http.ListenAndServe(*listenAddress, nil))
}
