package models

import (
	"database/sql"

	"github.com/uptrace/bun"
)

// +metric=slice
type PgStatActivity struct {
	bun.BaseModel   `bun:"pg_stat_activity"`
	DatID           int64          `bun:"datid" help:"OID of the database this backend is connected to" metric:"database_id,type:label"`
	DatName         string         `bun:"datname" help:"Name of the database this backend is connected to" metric:"database,type:label"`
	Usename         sql.NullString `bun:"usename" help:"Name of the user logged into this backend" metric:"username,type:label"`
	ApplicationName sql.NullString `bun:"application_name" help:"Name of the application that is connected to this backend" metric:",type:label"`
	ClientAddr      sql.NullString `bun:"client_addr" help:"IP address of the client connected to this backend" metric:",type:label"`
	State           sql.NullString `bun:"state" help:"Current overall state of this backend" metric:",type:label"`
	WaitEventType   sql.NullString `bun:"wait_event_type" help:"The type of event for which the backend is waiting, if any" metric:",type:label"`
	BackendType     sql.NullString `bun:"backend_type" help:"Type of current backend" metric:",type:label"`
	Connections     int64          `bun:"connections" help:"Number of active connections" metric:",type:gauge"`
}
