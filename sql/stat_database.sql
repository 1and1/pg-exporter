select current_setting('server_version_num')::int / 10000 >= 14 as is_pg_14;

\! curl -sf localhost:9135/metrics | grep "^postgresql_database_" | sed -e 's/database_id="[0-9]*"/database_id=""/' -e 's/ \([1-9]\|0\.\).*/ NNN/'
