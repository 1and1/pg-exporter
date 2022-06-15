package models

import (
	"github.com/uptrace/bun"
)

// +metric=slice
type PgStatStatements struct {
	bun.BaseModel `bun:"pg_stat_statements"`
	Datname       string           `bun:"datname" metric:"database,type:label"`
	Usename       string           `bun:"usename" metric:"username,type:label"`
	Query         string           `bun:"query" metric:"query,type:label"`
	Calls         int64            `bun:"calls" help:"Number of times executed" metric:"calls_total"`
	TotalTime     Milliseconds     `bun:"total_time" help:"Total time spent in the statement, in seconds" metric:"time_seconds_total"`
	MeanTime      NullMilliseconds `bun:"mean_time" help:"Mean time spent in the statement, in seconds" metric:"time_seconds_mean,type:gauge"`
}
