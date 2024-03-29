/*generated by binding gen*/
package models

import (
	"github.com/prometheus/client_golang/prometheus"
)

type PgFrozenXidSlice []PgFrozenXid

func (r PgFrozenXidSlice) ToMetrics(namespace string, subsystem string, ch chan<- prometheus.Metric, labelsKV ...string) error {
	for _, row := range []PgFrozenXid(r) {
		if err := row.ToMetrics(namespace, subsystem, ch, labelsKV...); err != nil {
			return err
		}
	}
	return nil
}

func (r *PgFrozenXid) ToMetrics(namespace string, subsystem string, ch chan<- prometheus.Metric, labelsKV ...string) error {
	labels := newLabels(labelsKV...)
	// labels
	labels["schema"] = r.Schema
	labels["table"] = r.TableName

	// optional labels

	// metrics
	// frozen_xid (CounterValue)
	frozenXid := float64(r.RelFrozenXid)
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, `frozen_xid`), `frozen XID for table`, nil, labels,
		), prometheus.CounterValue, frozenXid,
	)

	// optional metrics
	// toast_frozen_xid (CounterValue)
	if r.ToastFrozenXid.Valid {
		toastFrozenXid := float64(r.ToastFrozenXid.Int64)

		ch <- prometheus.MustNewConstMetric(
			prometheus.NewDesc(
				prometheus.BuildFQName(namespace, subsystem, `toast_frozen_xid`), `frozen XID for associated TOAST table`, nil, labels,
			), prometheus.CounterValue, toastFrozenXid,
		)
	}

	return nil
}
