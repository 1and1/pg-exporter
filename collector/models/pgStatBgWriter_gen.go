/*generated by binding gen*/
package models

import (
	"github.com/prometheus/client_golang/prometheus"
)

func (r *PgStatBgWriter) ToMetrics(namespace string, subsystem string, ch chan<- prometheus.Metric, labelsKV ...string) error {
	labels := newLabels(labelsKV...)
	// labels

	// optional labels

	// metrics
	// checkpoints_timed_total (CounterValue)
	checkpointsTimedTotal := float64(r.CheckpointsTimed)
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, `checkpoints_timed_total`), `Number of scheduled checkpoints that have been performed`, nil, labels,
		), prometheus.CounterValue, checkpointsTimedTotal,
	)

	// checkpoints_req_total (CounterValue)
	checkpointsReqTotal := float64(r.CheckpointsReq)
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, `checkpoints_req_total`), `Number of requested checkpoints that have been performed`, nil, labels,
		), prometheus.CounterValue, checkpointsReqTotal,
	)

	// checkpoint_write_seconds_total (CounterValue)

	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, `checkpoint_write_seconds_total`), `Total amount of time that has been spent in the portion of checkpoint processing where files are written to disk, in milliseconds`, nil, labels,
		), prometheus.CounterValue, r.CheckpointWriteTime,
	)

	// checkpoint_sync_seconds_total (CounterValue)

	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, `checkpoint_sync_seconds_total`), `Total amount of time that has been spent in the portion of checkpoint processing where files are synchronized to disk, in milliseconds`, nil, labels,
		), prometheus.CounterValue, r.CheckpointSyncTime,
	)

	// buffers_checkpoint_total (CounterValue)
	buffersCheckpointTotal := float64(r.BuffersCheckpoint)
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, `buffers_checkpoint_total`), `Number of buffers written during checkpoints`, nil, labels,
		), prometheus.CounterValue, buffersCheckpointTotal,
	)

	// buffers_clean_total (CounterValue)
	buffersCleanTotal := float64(r.BuffersClean)
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, `buffers_clean_total`), `Number of buffers written by the background writer`, nil, labels,
		), prometheus.CounterValue, buffersCleanTotal,
	)

	// maxwritten_clean_total (CounterValue)
	maxwrittenCleanTotal := float64(r.MaxwrittenClean)
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, `maxwritten_clean_total`), `Number of times the background writer stopped a cleaning scan because it had written too many buffers`, nil, labels,
		), prometheus.CounterValue, maxwrittenCleanTotal,
	)

	// buffers_backend_total (CounterValue)
	buffersBackendTotal := float64(r.BuffersBackend)
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, `buffers_backend_total`), `Number of buffers written directly by a backend`, nil, labels,
		), prometheus.CounterValue, buffersBackendTotal,
	)

	// buffers_backend_fsync_total (CounterValue)
	buffersBackendFsyncTotal := float64(r.BuffersBackendFsync)
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, `buffers_backend_fsync_total`), `Number of times a backend had to execute its own fsync call`, nil, labels,
		), prometheus.CounterValue, buffersBackendFsyncTotal,
	)

	// buffers_alloc_total (CounterValue)
	buffersAllocTotal := float64(r.BuffersAlloc)
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, `buffers_alloc_total`), `Number of buffers allocated`, nil, labels,
		), prometheus.CounterValue, buffersAllocTotal,
	)

	// stats_reset (CounterValue)
	var statsReset float64
	if r.StatsReset.IsZero() {
		statsReset = float64(0)
	} else {
		statsReset = float64(r.StatsReset.Unix())
	}
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, `stats_reset`), `Time at which these statistics were last reset`, nil, labels,
		), prometheus.CounterValue, statsReset,
	)

	// optional metrics

	return nil
}
