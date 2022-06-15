package models

import (
	"github.com/uptrace/bun"
)

// +metric=slice
type PgStatIOUserTables struct {
	bun.BaseModel `bun:"pg_statio_user_tables"`
	SchemaName    string `bun:"schemaname" help:"Name of the schema that this table is in" metric:"schema,type:label"`
	Relname       string `bun:"relname" help:"Name of this table" metric:"table,type:label"`
	HeapBlksRead  int64  `bun:"heap_blks_read" help:"Number of disk blocks read from this table" metric:"heap_blocks_read_total"`
	HeapBlksHit   int64  `bun:"heap_blks_hit" help:"Number of buffer hits in this table"  metric:"heap_blocks_hit_total"`
	IdxBlksRead   int64  `bun:"idx_blks_read" help:"Number of disk blocks read from all indexes on this table" metric:"index_blocks_read_total"`
	IdxBlksHit    int64  `bun:"idx_blks_hit" help:"Number of buffer hits in all indexes on this table" metric:"index_blocks_hit_total"`
	ToastBlksRead int64  `bun:"toast_blks_read" help:"Number of disk blocks read from this table's TOAST table" metric:"toast_blocks_read_total"`
	ToastBlksHit  int64  `bun:"toast_blks_hit" help:"Number of buffer hits in this table's TOAST table"  metric:"toast_blocks_hit_total"`
	TIDXBlksRead  int64  `bun:"tidx_blks_read" help:"Number of disk blocks read from this table's TOAST table indexes" metric:"toast_index_blocks_read_total"`
	TIDXBlksHit   int64  `bun:"tidx_blks_hit" help:"Number of buffer hits in this table's TOAST table indexes"  metric:"toast_index_blocks_hit_total"`
}
