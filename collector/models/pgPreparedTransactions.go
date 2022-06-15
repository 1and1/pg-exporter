package models

import (
	"time"

	"github.com/uptrace/bun"
)

// +metric=slice
type PgPreparedTransactions struct {
	bun.BaseModel `bun:"pg_prepared_xacts"`
	Oldest        time.Time `bun:"oldest" help:"Time at which the oldest transaction was prepared for commit" metric:",type:counter"`
	Database      string    `bun:"database" help:"Name of the database which the transactions where executed" metric:",type:label"`
	Count         int64     `bun:"count" help:"Number of prepared transactions" metric:"count,type:gauge"`
}
