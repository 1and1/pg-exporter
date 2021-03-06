/*generated by binding gen*/
package models

import (
	"github.com/prometheus/client_golang/prometheus"
)

type PgStatUserTablesSlice []PgStatUserTables

func (r PgStatUserTablesSlice) ToMetrics(namespace string, subsystem string, ch chan<- prometheus.Metric, labelsKV ...string) error {
	for _, row := range []PgStatUserTables(r) {
		if err := row.ToMetrics(namespace, subsystem, ch, labelsKV...); err != nil {
			return err
		}
	}
	return nil
}

func (r *PgStatUserTables) ToMetrics(namespace string, subsystem string, ch chan<- prometheus.Metric, labelsKV ...string) error {
	labels := newLabels(labelsKV...)
	// labels
	labels["schema"] = r.SchemaName
	labels["table"] = r.Relname

	// optional labels

	// metrics
	// seq_scans_total (CounterValue)
	seqScansTotal := float64(r.SeqScan)
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, `seq_scans_total`), `Number of sequential scans initiated on this table`, nil, labels,
		), prometheus.CounterValue, seqScansTotal,
	)

	// seq_rows_reads_total (CounterValue)
	seqRowsReadsTotal := float64(r.SeqTupRead)
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, `seq_rows_reads_total`), `Number of live rows fetched by sequential scans`, nil, labels,
		), prometheus.CounterValue, seqRowsReadsTotal,
	)

	// index_scans_total (CounterValue)
	indexScansTotal := float64(r.IdxScan)
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, `index_scans_total`), `Number of index scans initiated on this table`, nil, labels,
		), prometheus.CounterValue, indexScansTotal,
	)

	// index_rows_fetch_total (CounterValue)
	indexRowsFetchTotal := float64(r.IdxTupFetch)
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, `index_rows_fetch_total`), `Number of live rows fetched by index scans`, nil, labels,
		), prometheus.CounterValue, indexRowsFetchTotal,
	)

	// rows_inserted_total (CounterValue)
	rowsInsertedTotal := float64(r.TupIns)
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, `rows_inserted_total`), `Number of rows inserted`, nil, labels,
		), prometheus.CounterValue, rowsInsertedTotal,
	)

	// rows_updated_total (CounterValue)
	rowsUpdatedTotal := float64(r.TupUpd)
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, `rows_updated_total`), `Number of rows updated`, nil, labels,
		), prometheus.CounterValue, rowsUpdatedTotal,
	)

	// rows_deleted_total (CounterValue)
	rowsDeletedTotal := float64(r.TupDel)
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, `rows_deleted_total`), `Number of rows deleted`, nil, labels,
		), prometheus.CounterValue, rowsDeletedTotal,
	)

	// rows_hot_updated_total (CounterValue)
	rowsHotUpdatedTotal := float64(r.TupHotUpd)
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, `rows_hot_updated_total`), `Number of rows HOT updated`, nil, labels,
		), prometheus.CounterValue, rowsHotUpdatedTotal,
	)

	// live_rows (GaugeValue)
	liveRows := float64(r.LiveTup)
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, `live_rows`), `Estimated number of live rows`, nil, labels,
		), prometheus.GaugeValue, liveRows,
	)

	// dead_rows (GaugeValue)
	deadRows := float64(r.DeadTup)
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, `dead_rows`), `Estimated number of dead rows`, nil, labels,
		), prometheus.GaugeValue, deadRows,
	)

	// rows_modified_since_analyze_total (CounterValue)
	rowsModifiedSinceAnalyzeTotal := float64(r.ModSinceAnalyze)
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, `rows_modified_since_analyze_total`), `Estimated number of rows modified since this table was last analyzed`, nil, labels,
		), prometheus.CounterValue, rowsModifiedSinceAnalyzeTotal,
	)

	// last_vacuum (CounterValue)
	var lastVacuum float64
	if r.LastVacuum.IsZero() {
		lastVacuum = float64(0)
	} else {
		lastVacuum = float64(r.LastVacuum.Unix())
	}
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, `last_vacuum`), `Last time at which this table was manually vacuumed`, nil, labels,
		), prometheus.CounterValue, lastVacuum,
	)

	// last_autovacuum (CounterValue)
	var lastAutovacuum float64
	if r.LastAutoVacuum.IsZero() {
		lastAutovacuum = float64(0)
	} else {
		lastAutovacuum = float64(r.LastAutoVacuum.Unix())
	}
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, `last_autovacuum`), `Last time at which this table was vacuumed by the autovacuum daemon`, nil, labels,
		), prometheus.CounterValue, lastAutovacuum,
	)

	// last_analyze (CounterValue)
	var lastAnalyze float64
	if r.LastAnalyze.IsZero() {
		lastAnalyze = float64(0)
	} else {
		lastAnalyze = float64(r.LastAnalyze.Unix())
	}
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, `last_analyze`), `Last time at which this table was manually analyzed`, nil, labels,
		), prometheus.CounterValue, lastAnalyze,
	)

	// last_autoanalyze (CounterValue)
	var lastAutoanalyze float64
	if r.LastAutoAnalyze.IsZero() {
		lastAutoanalyze = float64(0)
	} else {
		lastAutoanalyze = float64(r.LastAutoAnalyze.Unix())
	}
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, `last_autoanalyze`), `Last time at which this table was analyzed by the autovacuum daemon`, nil, labels,
		), prometheus.CounterValue, lastAutoanalyze,
	)

	// vacuum_count (CounterValue)
	vacuumCount := float64(r.VacuumCount)
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, `vacuum_count`), `Number of times this table has been manually vacuumed`, nil, labels,
		), prometheus.CounterValue, vacuumCount,
	)

	// autovacuum_count (CounterValue)
	autovacuumCount := float64(r.AutoVacuumCount)
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, `autovacuum_count`), `Number of times this table has been vacuumed by the autovacuum daemon`, nil, labels,
		), prometheus.CounterValue, autovacuumCount,
	)

	// analyze_count (CounterValue)
	analyzeCount := float64(r.AnalyzeCount)
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, `analyze_count`), `Number of times this table has been manually analyzed`, nil, labels,
		), prometheus.CounterValue, analyzeCount,
	)

	// autoanalyze_count (CounterValue)
	autoanalyzeCount := float64(r.AutoAnalyzeCount)
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, `autoanalyze_count`), `Number of times this table has been analyzed by the autovacuum daemon`, nil, labels,
		), prometheus.CounterValue, autoanalyzeCount,
	)

	// optional metrics

	return nil
}
