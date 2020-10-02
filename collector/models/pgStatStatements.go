package models

// +metric=slice
type PgStatStatements struct {
    tableName struct{}         `sql:"pg_stat_statements"`
    Datname   string           `sql:"datname" metric:"database,type:label"`
    Usename   string           `sql:"usename" metric:"username,type:label"`
    Query     string           `sql:"query" metric:"query,type:label"`
    Calls     int64            `sql:"calls" help:"Number of times executed" metric:"calls_total"`
    TotalTime Milliseconds     `sql:"total_time" help:"Total time spent in the statement, in seconds" metric:"time_seconds_total"`
    MeanTime  NullMilliseconds `sql:"mean_time" help:"Mean time spent in the statement, in seconds" metric:"time_seconds_mean,type:gauge"`
}
