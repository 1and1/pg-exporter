package models

// +metric=slice
type PgStatStatements struct {
    tableName struct{}         `pg:"pg_stat_statements"`
    Datname   string           `pg:"datname" metric:"database,type:label"`
    Usename   string           `pg:"usename" metric:"username,type:label"`
    Query     string           `pg:"query" metric:"query,type:label"`
    Calls     int64            `pg:"calls" help:"Number of times executed" metric:"calls_total"`
    TotalTime Milliseconds     `pg:"total_time" help:"Total time spent in the statement, in seconds" metric:"time_seconds_total"`
    MeanTime  NullMilliseconds `pg:"mean_time" help:"Mean time spent in the statement, in seconds" metric:"time_seconds_mean,type:gauge"`
}
