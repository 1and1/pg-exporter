/*generated by binding gen*/
package models

import (
	"github.com/prometheus/client_golang/prometheus"
)

func (r *PgStatCheckpointer) ToMetrics(namespace string, subsystem string, ch chan<- prometheus.Metric, labelsKV ...string) error {
	labels := newLabels(labelsKV...)
	// labels

	// optional labels

	// metrics
	// num_timed_total (CounterValue)
	numTimedTotal := float64(r.NumTimed)
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, `num_timed_total`), `Number of Checkpoints  timed`, nil, labels,
		), prometheus.CounterValue, numTimedTotal,
	)

	// num_requested_total (CounterValue)
	numRequestedTotal := float64(r.NumRequested)
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, `num_requested_total`), `Number of Checkpoints requested`, nil, labels,
		), prometheus.CounterValue, numRequestedTotal,
	)

	// restartpoints_timed_total (CounterValue)
	restartpointsTimedTotal := float64(r.RestartpointsTimed)
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, `restartpoints_timed_total`), `Number of Restart Points timed`, nil, labels,
		), prometheus.CounterValue, restartpointsTimedTotal,
	)

	// restartpoints_req_total (CounterValue)
	restartpointsReqTotal := float64(r.RestartpointsReq)
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, `restartpoints_req_total`), `Number of Restart Points requested`, nil, labels,
		), prometheus.CounterValue, restartpointsReqTotal,
	)

	// restartpoints_done_total (CounterValue)
	restartpointsDoneTotal := float64(r.RestartpointsDone)
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, `restartpoints_done_total`), `Number of Restart Points done`, nil, labels,
		), prometheus.CounterValue, restartpointsDoneTotal,
	)

	// ckp_write_time_total (CounterValue)

	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, `ckp_write_time_total`), `Checkpoint write Time`, nil, labels,
		), prometheus.CounterValue, r.CkpWriteTime,
	)

	// ckp_sync_time_total (CounterValue)

	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, `ckp_sync_time_total`), `Checkpoint sync Time`, nil, labels,
		), prometheus.CounterValue, r.CkpSyncTime,
	)

	// ckp_buffers_written_total (CounterValue)
	ckpBuffersWrittenTotal := float64(r.BuffersWritten)
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, `ckp_buffers_written_total`), `Checkpoint Buffers written`, nil, labels,
		), prometheus.CounterValue, ckpBuffersWrittenTotal,
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
