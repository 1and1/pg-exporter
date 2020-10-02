package generator

import (
	"fmt"
	"io"

	"k8s.io/gengo/generator"
	"k8s.io/gengo/namer"
	"k8s.io/gengo/types"
)

func DefaultNameSystem() string {
	return "public"
}

func NameSystems() namer.NameSystems {
	return namer.NameSystems{
		"public":  namer.NewPublicNamer(0),
		"private": namer.NewPrivateNamer(0),
		"raw":     namer.NewRawNamer("", nil),
	}
}

type mappingGen struct {
	generator.DefaultGen
	outputPackage string
	typeToMatch   *types.Type
	imports       namer.ImportTracker
	isSlice       bool
}

func (m *mappingGen) Name() string {
	return "mapping generator"
}

func (m *mappingGen) Filter(c *generator.Context, t *types.Type) bool {
	if t == m.typeToMatch {
		_, m.isSlice = extractTag("metric", t.CommentLines)
		return true
	}
	return false
}

func (m *mappingGen) Namers(*generator.Context) namer.NameSystems {
	return namer.NameSystems{
		"raw": namer.NewRawNamer(m.outputPackage, m.imports),
	}
}

func (m *mappingGen) Init(*generator.Context, io.Writer) error {
	return nil
}

func (m *mappingGen) Finalize(*generator.Context, io.Writer) error {
	return nil
}

func (m *mappingGen) PackageVars(*generator.Context) []string {
	return nil
}

func (m *mappingGen) PackageConsts(*generator.Context) []string {
	return nil
}

func (m *mappingGen) GenerateType(c *generator.Context, t *types.Type, w io.Writer) error {
	sw := generator.NewSnippetWriter(w, c, "$", "$")

	// get the model description
	desc := parseModel(t)

	// create the slice handler if required
	if m.isSlice {
		sw.Do(sliceType+"\n", t).
			Do(sliceToMetricsFunc+"\n", t)
	}

	// create the core ToMetrics func
	sw.Do(toMetricsFunc, t)

	// generate the labels list
	sw.Do("// labels\n", nil)
	for _, label := range desc.labels {
		var tmpl string
		switch label.T.Name.String() {
		case "int":
			tmpl = intLabel
		case "int64":
			tmpl = int64Label
		case "string":
			tmpl = stringLabel
		}
		tmpl += "\n"
		sw.Do(tmpl, label)
	}
	sw.Do("\n", nil)

	// generate the optional labels
	sw.Do("// optional labels\n", nil)
	for _, label := range desc.optionalLabels {
		var tmpl string
		switch label.T.Name.String() {
		case "database/sql.NullString":
			tmpl = nullStringLabel
		case "database/sql.NullBool":
			tmpl = nullBoolLabel
		}
		tmpl += "\n"
		sw.Do(tmpl, label)
	}
	sw.Do("\n", nil)

	// create the metrics
	sw.Do("// metrics \n", nil)
	for _, metric := range desc.metrics {
		sw.Do(`// $.Name$ ($.PromType$)`+"\n", metric)
		var tmpl string
		switch metric.T.Name.String() {
		case "int64", "int":
			tmpl = int64Metric
		case "time.Time":
			tmpl = timeMetric
		case "float64":
			// no convertion needed, we ust replace the VarName
			metric.VarName = fmt.Sprintf("r.%s", metric.Field)
		case "./collector/models.Milliseconds", "github.com/1and1/pg-exporter/collector/models.Milliseconds":
			tmpl = millisecondsMetric
		}
		tmpl += "\n"
		sw.Do(tmpl, metric)
		sw.Do(sendMetric+"\n\n", metric)
	}
	sw.Do("\n", nil)

	// create the optional metrics
	sw.Do("// optional metrics \n", nil)
	for _, metric := range desc.optionalMetrics {
		sw.Do(`// $.Name$ ($.PromType$)`+"\n", metric)
		var tmpl string
		switch metric.T.Name.String() {
		case "database/sql.NullInt64":
			tmpl = nullInt64Metric
		case "database/sql.NullFloat64":
			tmpl = nullFloat64Metric
		case "./collector/models.NullMilliseconds", "github.com/1and1/pg-exporter/collector/models.NullMilliseconds":
			tmpl = nullMillisecondsMetric
		}
		tmpl += "\n"
		sw.Do(tmpl, metric)
		sw.Do(sendMetric+"\n", metric)
		sw.Do("}\n", nil)
	}
	sw.Do("\n", nil)

	sw.Do("return nil", nil)
	sw.Do("}\n", t)

	return sw.Error()
}

func (m *mappingGen) Imports(*generator.Context) []string {
	return append(m.imports.ImportLines(),
		"database/sql",
		"strconv",
		"github.com/prometheus/client_golang/prometheus",
	)
}
