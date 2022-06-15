package models

import (
	"time"

	"github.com/uptrace/bun"
)

// +metric=slice
type PgStatUserTables struct {
	bun.BaseModel    `bun:"pg_stat_user_tables"`
	SchemaName       string    `bun:"schemaname" help:"Name of the schema that this table is in" metric:"schema,type:label"`
	Relname          string    `bun:"relname" help:"Name of this table" metric:"table,type:label"`
	SeqScan          int64     `bun:"seq_scan" help:"Number of sequential scans initiated on this table" metric:"seq_scans_total"`
	SeqTupRead       int64     `bun:"seq_tup_read" help:"Number of live rows fetched by sequential scans" metric:"seq_rows_reads_total"`
	IdxScan          int64     `bun:"idx_scan" help:"Number of index scans initiated on this table" metric:"index_scans_total"`
	IdxTupFetch      int64     `bun:"idx_tup_fetch" help:"Number of live rows fetched by index scans" metric:"index_rows_fetch_total"`
	TupIns           int64     `bun:"n_tup_ins" help:"Number of rows inserted" metric:"rows_inserted_total"`
	TupUpd           int64     `bun:"n_tup_upd" help:"Number of rows updated" metric:"rows_updated_total"`
	TupDel           int64     `bun:"n_tup_del" help:"Number of rows deleted" metric:"rows_deleted_total"`
	TupHotUpd        int64     `bun:"n_tup_hot_upd" help:"Number of rows HOT updated" metric:"rows_hot_updated_total"`
	LiveTup          int64     `bun:"n_live_tup" help:"Estimated number of live rows" metric:"live_rows,type:gauge"`
	DeadTup          int64     `bun:"n_dead_tup" help:"Estimated number of dead rows" metric:"dead_rows,type:gauge"`
	ModSinceAnalyze  int64     `bun:"n_mod_since_analyze" help:"Estimated number of rows modified since this table was last analyzed" metric:"rows_modified_since_analyze_total"`
	LastVacuum       time.Time `bun:"last_vacuum" help:"Last time at which this table was manually vacuumed"`
	LastAutoVacuum   time.Time `bun:"last_autovacuum" help:"Last time at which this table was vacuumed by the autovacuum daemon"`
	LastAnalyze      time.Time `bun:"last_analyze" help:"Last time at which this table was manually analyzed"`
	LastAutoAnalyze  time.Time `bun:"last_autoanalyze" help:"Last time at which this table was analyzed by the autovacuum daemon"`
	VacuumCount      int64     `bun:"vacuum_count" help:"Number of times this table has been manually vacuumed"`
	AutoVacuumCount  int64     `bun:"autovacuum_count" help:"Number of times this table has been vacuumed by the autovacuum daemon"`
	AnalyzeCount     int64     `bun:"analyze_count" help:"Number of times this table has been manually analyzed"`
	AutoAnalyzeCount int64     `bun:"autoanalyze_count" help:"Number of times this table has been analyzed by the autovacuum daemon"`
}
