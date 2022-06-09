package models

// +metric=slice
type PgDatabase struct {
	tableName struct{} `pg:"pg_database"`
	Datname   string   `pg:"datname" metric:"database,type:label"`
	FrozenXID int64    `pg:"datfrozenxid" help:"All transaction IDs before this one have been replaced with a permanent transaction ID in this database" metric:"frozen_xid"`
	MinMXID   int64    `pg:"datminmxid" help:"All multixact IDs before this one have been replaced with a transaction ID in this database"  metric:"min_mxid"`
}
