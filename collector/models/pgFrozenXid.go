package models

import (
	"database/sql"

	"github.com/uptrace/bun"
)

// +metric=slice
type PgFrozenXid struct {
	bun.BaseModel  `bun:"pg_prepared_xacts"`
	Schema         string        `bun:"schema" metric:"schema,type:label"`
	TableName      string        `bun:"table_name" metric:"table,type:label"`
	RelFrozenXid   int64         `bun:"relfrozenxid" help:"frozen XID for table" metric:"frozen_xid"`
	ToastFrozenXid sql.NullInt64 `bun:"toastfrozenxid" help:"frozen XID for associated TOAST table" metric:"toast_frozen_xid"`
}
