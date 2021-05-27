package collector

import (
	"context"
	"regexp"
	"strconv"
	"strings"

	"github.com/go-pg/pg/v9"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	// subsystem
	settings = "settings"
)

type pgSetting struct {
	tableName struct{} `sql:"pg_settings"`
	Name      string   `sql:"name"`
	Setting   string   `sql:"setting"`
	ShortDesc string   `sql:"short_desc"`
	Unit      string   `sql:"unit"`
	Vartype   string   `sql:"vartype"`
	MinVal    string   `sql:"min_val"`
	MaxVal    string   `sql:"max_val"`
	Enumvals  []string `sql:"enumvals,array"`
}

// ScrapeSettings scrapes from pg_settings
type ScrapeSettings struct{}

// Name of the Scraper
func (ScrapeSettings) Name() string {
	return "pg_settings"
}

// Help describes the role of the Scraper
func (ScrapeSettings) Help() string {
	return "Collect from pg_settings"
}

// Version returns minimum PostgreSQL version
func (ScrapeSettings) Version() int {
	return 0
}

// Type returns the scrape type
func (ScrapeSettings) Type() ScrapeType {
	return SCRAPEGLOBAL
}

// Scrape collects data from database connection and sends it over channel as prometheus metric.
func (ScrapeSettings) Scrape(ctx context.Context, db *pg.DB, ch chan<- prometheus.Metric) error {

	var settingsRes []pgSetting
	if err := db.ModelContext(ctx, &settingsRes).WhereIn("vartype IN (?)", []string{"bool", "integer", "real"}).Select(); err != nil {
		return err
	}

	unitRE := regexp.MustCompile(`(?P<mult>\d+)?(?P<unit>.*)`)

	// as we have to check very returned row and convert it to a metric on a different way
	// we can not use the generic helper here
	for _, setting := range settingsRes {
		// normalize the settings name
		setting.Name = strings.ReplaceAll(setting.Name, ".", "_")
		value := 1.0
		switch setting.Vartype {
		case "bool":
			setting.ShortDesc += " (bool)"
			switch setting.Setting {
			case "off":
				value = 0.0
			case "on":
				value = 1.0
			}
		case "integer":
			intValue, err := strconv.Atoi(setting.Setting)
			if err != nil {
				return err
			}
			value = float64(intValue)
		case "real":
			realValue, err := strconv.ParseFloat(setting.Setting, 64)
			if err != nil {
				return err
			}
			value = realValue
		}

		// get the pre/suffix if any from the value

		unitCaptures := unitRE.FindStringSubmatch(setting.Unit)

		// check if we have an unit type
		switch unitCaptures[2] {
		case "min":
			// convert minute to seconds
			value = value * 60
		case "ms":
			// convert milliseconds to seconds
			value = value / 1000
		case "kB":
			// convert kilobytes to bytes
			value = value * 1024
		case "MB":
			// convert megabytes to bytes
			value = value * 1024 * 1024
		}
		if unitCaptures[1] != "" {
			multiplicator, err := strconv.Atoi(unitCaptures[1])
			if err != nil {
				return err
			}
			value = value * float64(multiplicator)
		}
		// create a metric from this
		if setting.MinVal != "" {
			setting.ShortDesc += " min=" + setting.MinVal
		}
		if setting.MaxVal != "" {
			setting.ShortDesc += " max=" + setting.MinVal
		}

		ch <- prometheus.MustNewConstMetric(
			prometheus.NewDesc(
				prometheus.BuildFQName(namespace, settings, setting.Name), setting.ShortDesc, nil, nil,
			), prometheus.GaugeValue, value,
		)
	}
	return nil
}
