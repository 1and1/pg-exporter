package models

// +metric=slice
type PgDatabase struct {
	tableName struct{} `sql:"pg_database"`
	Datname   string   `sql:"datname" metric:"database,type:label"`
	FrozenXID int64    `sql:"datfrozenxid" help:"All transaction IDs before this one have been replaced with a permanent transaction ID in this database" metric:"frozen_xid"`
	MinMXID   int64    `sql:"datminmxid" help:"All multixact IDs before this one have been replaced with a transaction ID in this database"  metric:"min_mxid"`
}
