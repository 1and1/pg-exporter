package models

// +metric=slice
type PgStatIOUserTables struct {
	tableName     struct{} `sql:"pg_statio_user_tables"`
	SchemaName    string   `sql:"schemaname" help:"Name of the schema that this table is in" metric:"schema,type:label"`
	Relname       string   `sql:"relname" help:"Name of this table" metric:"table,type:label"`
	HeapBlksRead  int64    `sql:"heap_blks_read" help:"Number of disk blocks read from this table" metric:"heap_blocks_read_total"`
	HeapBlksHit   int64    `sql:"heap_blks_hit" help:"Number of buffer hits in this table"  metric:"heap_blocks_hit_total"`
	IdxBlksRead   int64    `sql:"idx_blks_read" help:"Number of disk blocks read from all indexes on this table" metric:"index_blocks_read_total"`
	IdxBlksHit    int64    `sql:"idx_blks_hit" help:"Number of buffer hits in all indexes on this table" metric:"index_blocks_hit_total"`
	ToastBlksRead int64    `sql:"toast_blks_read" help:"Number of disk blocks read from this table's TOAST table" metric:"toast_blocks_read_total"`
	ToastBlksHit  int64    `sql:"toast_blks_hit" help:"Number of buffer hits in this table's TOAST table"  metric:"toast_blocks_hit_total"`
	TIDXBlksRead  int64    `sql:"tidx_blks_read" help:"Number of disk blocks read from this table's TOAST table indexes" metric:"toast_index_blocks_read_total"`
	TIDXBlksHit   int64    `sql:"tidx_blks_hit" help:"Number of buffer hits in this table's TOAST table indexes"  metric:"toast_index_blocks_hit_total"`
}
