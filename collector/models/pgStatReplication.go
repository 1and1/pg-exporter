package models

import (
	"database/sql"
)

// +metric=slice
type PgStatReplication struct {
	tableName       struct{}       `pg:"pg_stat_replication,discard_unknown_columns"`
	PID             int            `pg:"pid" help:"Process ID of a WAL sender process" metric:",type:label"`
	ApplicationName string         `pg:"application_name" help:"Name of the application that is connected to this WAL sender" metric:",type:label"`
	ClientAddr      sql.NullString `pg:"client_addr" help:"IP address of the client connected to this WAL sender." metric:",type:label"` // TODO: should this be optional?
	BackendXmin     int64          `pg:"backend_xmin" help:"This standby's xmin horizon reported by hot_standby_feedback."`
	SentBytesLag    int64          `pg:"sent_lag_bytes" help:"Number of bytes not yet sent on this connection" metric:",type:gauge"`
	WriteBytesLag   int64          `pg:"write_lag_bytes" help:"Number of bytes not yet written to this on this standby server" metric:",type:gauge"`
	FlushBytesLag   int64          `pg:"flush_lag_bytes" help:"Number of bytes not yet flushed to disk by this standby server" metric:",type:gauge"`
	ReplayBytesLag  int64          `pg:"replay_lag_bytes" help:"Number of bytes not yet replayed on this standby server" metric:",type:gauge"`

	// PostgreSQL 10 and newer
	WriteLag  sql.NullFloat64 `pg:"write_lag" help:"Time elapsed between flushing recent WAL locally and receiving notification that this standby server has written it"  metric:"write_lag_seconds,type:gauge"`
	FlushLag  sql.NullFloat64 `pg:"flush_lag" help:"Time elapsed between flushing recent WAL locally and receiving notification that this standby server has written and flushed it"  metric:"flush_lag_seconds,type:gauge"`
	ReplayLag sql.NullFloat64 `pg:"replay_lag" help:"Time elapsed between flushing recent WAL locally and receiving notification that this standby server has written, flushed and applied it"  metric:"replay_lag_seconds,type:gauge"`
}
