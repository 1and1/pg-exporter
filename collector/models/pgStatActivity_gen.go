/*generated by binding gen*/
package models

import (
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
)

type PgStatActivitySlice []PgStatActivity

func (r PgStatActivitySlice) ToMetrics(namespace string, subsystem string, ch chan<- prometheus.Metric, labelsKV ...string) error {
	for _, row := range []PgStatActivity(r) {
		if err := row.ToMetrics(namespace, subsystem, ch, labelsKV...); err != nil {
			return err
		}
	}
	return nil
}

func (r *PgStatActivity) ToMetrics(namespace string, subsystem string, ch chan<- prometheus.Metric, labelsKV ...string) error {
	labels := newLabels(labelsKV...)
	// labels
	labels["database_id"] = strconv.FormatInt(r.DatID, 10)
	labels["database"] = r.DatName

	// optional labels
	if r.Usename.Valid {
		labels["username"] = r.Usename.String
	}
	if r.ApplicationName.Valid {
		labels["application_name"] = r.ApplicationName.String
	}
	if r.ClientAddr.Valid {
		labels["client_addr"] = r.ClientAddr.String
	}
	if r.State.Valid {
		labels["state"] = r.State.String
	}
	if r.WaitEventType.Valid {
		labels["wait_event_type"] = r.WaitEventType.String
	}
	if r.BackendType.Valid {
		labels["backend_type"] = r.BackendType.String
	}

	// metrics
	// connections (GaugeValue)
	connections := float64(r.Connections)
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, `connections`), `Number of active connections`, nil, labels,
		), prometheus.GaugeValue, connections,
	)

	// optional metrics

	return nil
}
