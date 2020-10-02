package generator

import (
    "reflect"

    "github.com/iancoleman/strcase"
    "github.com/vmihailenco/tagparser"
    "k8s.io/gengo/types"
    "k8s.io/klog"
)

type modelDescription struct {
    labels          []modelField
    optionalLabels  []modelField
    metrics         []modelField
    optionalMetrics []modelField
}

type modelField struct {
    Field    string
    Name     string
    VarName  string
    T        *types.Type
    Help     string
    PromType string
}

func parseModel(t *types.Type) *modelDescription {
    desc := &modelDescription{}

    for _, member := range t.Members {
        if member.Tags == "" {
            continue
        }
        tag := reflect.StructTag(member.Tags)
        helpTag := tag.Get("help")
        metricTag := tagparser.Parse(tag.Get("metric"))
        sqlTag := tagparser.Parse(tag.Get("sql"))

        // check if we have a metric tag, if not relay on the sql name
        var metricName string
        // we skip metric tags witht `-` as value
        if metricTag.Name == "-" {
            continue
        }

        if metricTag.Name != "" {
            metricName = metricTag.Name
        } else if sqlTag.Name != "" {
            metricName = sqlTag.Name
        } else {
            // we skip every column where we can't find a valid name
            continue
        }

        field := modelField{
            Field:   member.Name,
            Name:    metricName,
            T:       member.Type,
            Help:    helpTag,
            VarName: strcase.ToLowerCamel(metricName),
        }

        isLabel := false
        switch metricTag.Options["type"] {
        case "counter", "":
            field.PromType = "CounterValue"
        case "gauge":
            field.PromType = "GaugeValue"
        case "label":
            isLabel = true
        }

        if helpTag == "" && !isLabel {
            continue
        }

        if isLabel {
            switch field.T.Name.String() {
            case "int", "int64", "string":
                desc.labels = append(desc.labels, field)
            case "database/sql.NullString", "database/sql.NullBool":
                desc.optionalLabels = append(desc.optionalLabels, field)
            default:
                klog.Fatalf("type for label %v unknown: %v", field.Field, field.T.Name.String())
            }
            continue
        }

        switch field.T.Name.String() {
        case "int64", "int", "time.Time", "float64",
            "./collector/models.Milliseconds", "github.com/1and1/pg-exporter/collector/models.Milliseconds":
            desc.metrics = append(desc.metrics, field)
        case "database/sql.NullInt64", "database/sql.NullFloat64",
            "./collector/models.NullMilliseconds", "github.com/1and1/pg-exporter/collector/models.NullMilliseconds":
            desc.optionalMetrics = append(desc.optionalMetrics, field)
        default:
            klog.Fatalf("type for metric %v unknown: %v", field.Field, field.T.Name.String())
        }
    }

    return desc
}
