package generator

// slice type definition
const sliceType = `type $.|public$Slice []$.|public$`

// slice ToMetrics definition
const sliceToMetricsFunc = `func (r $.|public$Slice) ToMetrics(namespace string, subsystem string, ch chan<- prometheus.Metric, labelsKV ...string) error {
	for _, row := range []$.|public$(r) {
        if err := row.ToMetrics(namespace,subsystem,ch,labelsKV...); err != nil {
            return err
        }
    }
    return nil
}
`

// function header for ToMetrics mapper
const toMetricsFunc = `func (r *$.|public$) ToMetrics(namespace string, subsystem string, ch chan<- prometheus.Metric, labelsKV ...string) error {
	labels := newLabels(labelsKV...)
`

// float64 metric
const float64Metric = `$.VarName$ := r.$.Field$`

// time Metric
const timeMetric = `var $.VarName$ float64
			if r.$.Field$.IsZero() {
				$.VarName$ = float64(0)
			} else {
				$.VarName$ = float64(r.$.Field$.Unix())
		}`

// int64 metric
const int64Metric = `$.VarName$ := float64(r.$.Field$)`

// milliseconds metric
const millisecondsMetric = `$.VarName$ := r.$.Field$.Seconds()`

// NullMilliseconds metric
const nullMillisecondsMetric = `if r.$.Field$.Valid {
  $.VarName$ := r.$.Field$.Seconds()
`

const nullTimeMetric = `if !r.$.Field$.IsZero() {
  $.VarName$ := float64(r.$.Field$.Unix())
`

// sql.NullInt64 metric
const nullInt64Metric = `if r.$.Field$.Valid {
$.VarName$ := float64(r.$.Field$.Int64)
`

// sql.NullInt64 metric
const nullFloat64Metric = `if r.$.Field$.Valid {
$.VarName$ := r.$.Field$.Float64
`

// int label
const intLabel = `labels["$.Name$"] = strconv.Itoa(r.$.Field$)`

// int64 label
const int64Label = `labels["$.Name$"] = strconv.FormatInt(r.$.Field$, 10)`

// string label
const stringLabel = `labels["$.Name$"] = r.$.Field$`

// sql.NullString label
const nullStringLabel = `if r.$.Field$.Valid {
			labels["$.Name$"] = r.$.Field$.String
		}`

// sql.NullBool label
const nullBoolLabel = `if r.$.Field$.Valid {
			if r.$.Field$.Bool {
				labels["$.Name$"] = "true"
			} else {
				labels["$.Name$"] = "false"
			}
		}`

// send a metric
const sendMetric = `ch <- prometheus.MustNewConstMetric(
	prometheus.NewDesc(
				prometheus.BuildFQName(namespace, subsystem, ` + "`$.Name$`), `$.Help$`" + `, nil, labels,
	), prometheus.$.PromType$, $.VarName$,
)`
