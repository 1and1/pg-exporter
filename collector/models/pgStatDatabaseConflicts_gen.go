/*generated by binding gen*/
package models

import (
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
)

type PgStatDatabaseConflictsSlice []PgStatDatabaseConflicts

func (r PgStatDatabaseConflictsSlice) ToMetrics(namespace string, subsystem string, ch chan<- prometheus.Metric, labelsKV ...string) error {
	for _, row := range []PgStatDatabaseConflicts(r) {
		if err := row.ToMetrics(namespace, subsystem, ch, labelsKV...); err != nil {
			return err
		}
	}
	return nil
}

func (r *PgStatDatabaseConflicts) ToMetrics(namespace string, subsystem string, ch chan<- prometheus.Metric, labelsKV ...string) error {
	labels := newLabels(labelsKV...)
	// labels
	labels["database_id"] = strconv.FormatInt(r.DatID, 10)
	labels["database"] = r.DatName

	// optional labels

	// metrics
	// confl_tablespace_total (CounterValue)
	conflTablespaceTotal := float64(r.ConfTablespace)
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, `confl_tablespace_total`), `Number of queries in this database that have been canceled due to dropped tablespaces`, nil, labels,
		), prometheus.CounterValue, conflTablespaceTotal,
	)

	// confl_lock_total (CounterValue)
	conflLockTotal := float64(r.ConfLock)
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, `confl_lock_total`), `Number of queries in this database that have been canceled due to lock timeouts`, nil, labels,
		), prometheus.CounterValue, conflLockTotal,
	)

	// confl_snapshot_total (CounterValue)
	conflSnapshotTotal := float64(r.ConfSnapshot)
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, `confl_snapshot_total`), `Number of queries in this database that have been canceled due to old snapshots`, nil, labels,
		), prometheus.CounterValue, conflSnapshotTotal,
	)

	// confl_bufferpin_total (CounterValue)
	conflBufferpinTotal := float64(r.ConfBufferpin)
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, `confl_bufferpin_total`), `Number of queries in this database that have been canceled due to pinned buffers`, nil, labels,
		), prometheus.CounterValue, conflBufferpinTotal,
	)

	// confl_deadlock_total (CounterValue)
	conflDeadlockTotal := float64(r.ConfDeadlock)
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, `confl_deadlock_total`), `Number of queries in this database that have been canceled due to deadlocks`, nil, labels,
		), prometheus.CounterValue, conflDeadlockTotal,
	)

	// optional metrics

	return nil
}
