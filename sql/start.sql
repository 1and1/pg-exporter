show port \gset
\setenv PGPORT :port
show unix_socket_directories \gset
\setenv PGHOST :unix_socket_directories
select current_database() as database \gset
\setenv PGDATABASE :database

\! ./bin/pg_exporter --log.level=warn & echo $! > pid; sleep 0.1

\! curl -sf localhost:9135/metrics | grep postgresql_up
