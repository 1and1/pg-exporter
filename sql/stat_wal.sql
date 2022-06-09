select current_setting('server_version_num')::int / 10000 >= 14 as is_pg_14;
show fsync; -- needs to be on for postgresql_wal_sync_count
select pg_stat_reset_shared('wal');

\! curl -sf localhost:9135/metrics | grep "^postgresql_wal_" | sed -e 's/ \([1-9]\|0\.\).*/ NNN/'
