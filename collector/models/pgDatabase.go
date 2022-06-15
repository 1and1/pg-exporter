package models

import (
	"github.com/uptrace/bun"
)

// +metric=slice
type PgDatabase struct {
	bun.BaseModel `bun:"pg_database"`
	Datname       string `bun:"datname" metric:"database,type:label"`
	FrozenXID     int64  `bun:"datfrozenxid" help:"All transaction IDs before this one have been replaced with a permanent transaction ID in this database" metric:"frozen_xid"`
	MinMXID       int64  `bun:"datminmxid" help:"All multixact IDs before this one have been replaced with a transaction ID in this database"  metric:"min_mxid"`
}
