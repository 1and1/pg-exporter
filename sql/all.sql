select current_setting('server_version_num')::int / 10000 as pg_major_version;

-- remove all values from the stats
-- remove database ids
-- remove Go and PostgreSQL version numbers
-- remove postgresql_locks_count{mode="RowExclusiveLock"} line which sometimes pops up
-- remove debug_discard_caches help since its max value depends on --enable-cassert
\! curl -sf localhost:9135/metrics | sed -e 's/^\([a-z].*\) .*/\1/' -e 's/database_id="[0-9]*"/database_id=""/' -e 's/version="[^"]*"/version=""/' -e '/mode="RowExclusiveLock"/d' -e '/# HELP postgresql_settings_debug_discard_caches/d'
