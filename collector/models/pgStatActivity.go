package models

import (
	"database/sql"
)

// +metric=slice
type PgStatActivity struct {
	tableName       struct{}       `sql:"pg_stat_activity"`
	DatID           int64          `sql:"datid" help:"OID of the database this backend is connected to" metric:"database_id,type:label"`
	DatName         string         `sql:"datname" help:"Name of the database this backend is connected to" metric:"database,type:label"`
	Usename         sql.NullString `sql:"usename" help:"Name of the user logged into this backend" metric:"username,type:label"`
	ApplicationName sql.NullString `sql:"application_name" help:"Name of the application that is connected to this backend" metric:",type:label"`
	ClientAddr      sql.NullString `sql:"client_addr" help:"IP address of the client connected to this backend" metric:",type:label"`
	State           sql.NullString `sql:"state" help:"Current overall state of this backend" metric:",type:label"`
	Connections     int64          `sql:"connections" help:"Number of active connections" metric:",type:gauge"`
}
