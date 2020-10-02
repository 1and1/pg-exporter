package models

import (
	"time"
)

// +metric=slice
type PgPreparedTransactions struct {
	tableName struct{}  `sql:"pg_prepared_xacts"`
	Oldest    time.Time `sql:"oldest" help:"Time at which the oldest transaction was prepared for commit" metric:",type:counter"`
	Database  string    `sql:"database" help:"Name of the database which the transactions where executed" metric:",type:label"`
	Count     int64     `sql:"count" help:"Number of prepared transactions" metric:"count,type:gauge"`
}
