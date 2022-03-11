package models

import (
	"time"
)

// +metric=slice
type PgStatReplicationSlots struct {
	tableName   struct{}  `pg:"pg_stat_replication_slots"`
	SlotName    string    `pg:"slot_name" help:"A unique, cluster-wide identifier for the replication slot" metric:"slot_name,type:label"`
	SpillTxns   int64     `pg:"spill_txns" help:"Number of transactions spilled to disk once the memory used by logical decoding to decode changes from WAL has exceeded logical_decoding_work_mem"`
	SpillCount  int64     `pg:"spill_count" help:"Number of times transactions were spilled to disk while decoding changes from WAL for this slot"`
	SpillBytes  int64     `pg:"spill_bytes" help:"Amount of decoded transaction data spilled to disk while performing decoding of changes from WAL for this slot"`
	StreamTxns  int64     `pg:"stream_txns" help:"Number of in-progress transactions streamed to the decoding output plugin after the memory used by logical decoding to decode changes from WAL for this slot has exceeded logical_decoding_work_mem"`
	StreamCount int64     `pg:"stream_count" help:"Number of times in-progress transactions were streamed to the decoding output plugin while decoding changes from WAL for this slot"`
	StreamBytes int64     `pg:"stream_bytes" help:"Amount of transaction data decoded for streaming in-progress transactions to the decoding output plugin while decoding changes from WAL for this slot"`
	TotalTxns   int64     `pg:"total_txns" help:"Number of decoded transactions sent to the decoding output plugin for this slot"`
	TotalBytes  int64     `pg:"total_bytes" help:"Amount of transaction data decoded for sending transactions to the decoding output plugin while decoding changes from WAL for this slot"`
	StatsReset  time.Time `pg:"stats_reset" help:"Time at which these statistics were last reset"`
}
