# PostgreSQL Server Exporter

Prometheus exporter for PostgreSQL server metrics.

Supported version:
- PostgreSQL >= 9.6

## Project state

This project is in early beta. There is missing documentation,
things may still change and the default scrapers are not final, and could change
at any time.


## Building and running

### Build
```shell
make
```

### Running

```shell
export PGHOST="<host>"
export PGUSER="<user>"
# ....
./bin/pg_exporter <flags>
```

#### Connection URL

`pg_exporter` supports `libpq` parsable connection strings and environments variables. If you want
to use a connection string, it could be provided via the `DATA_SOURCE_NAME` variable

### Collector Flags

|                 Name                       | PostgreSQL version  |                 Description                    |
|--------------------------------------------|---------------------|------------------------------------------------|
| collect.include.database                   | all                 | Database to include in the collection. If defined only the given databases are scraped. Can be defined multiple times |
| collect.exclude.database                   | all                 | Database to exclude from the collection. Only used if collect.include.database is not given. Can be defined multiple times |
| collect.info                               | all                 | Collect postgresql information                 |
| collect.pg_stat_activity                   | all                 | Collect from pg_stat_activity                  |
| collect.pg_stat_activity.with_username     | all                 | Include username in session statistics         |
| collect.pg_stat_activity.with_appname      | all                 | Include application name in session statistics |
| collect.pg_stat_activity.with_clientaddr   | all                 | Include application name in session statistics |
| collect.pg_stat_activity.with_state        | all                 | Include session state in session statistics    |
| collect.pg_stat_activity.with_wait_type    | all                 | Include wait_event_type in session statistics  |
| collect.pg_stat_activity.with_backend_type | >= 11               | Include backend_type in session statistics     |
| collect.pg_stat_user_tables                | all                 | Collect from pg_stat_user_tables               |
| collect.pg_settings                        | all                 | Collect from pg_settings                       |
| collect.pg_locks                           | all                 | Collect from pg_locks                          |
| collect.pg_stat_bgwriter                   | all                 | Collect from pg_stat_bgwriter                  |
| collect.pg_stat_database                   | all                 | Collect from pg_stat_database                  |
| collect.pg_stat_database_conflicts         | all                 | Collect from pg_stat_database_conflicts        |
| collect.pg_stat_archiver                   | all                 | Collect from pg_stat_statements                |
| collect.pg_stat_statements                 | all                 | Collect from pg_stat_archiver                  |
| collect.pg_stat_replication                | all                 | Collect from pg_stat_replication               |
| collect.pg_prepared_xacts                  | all                 | Collect from pg_prepared_xacts                 |
| collect.pg_statio_user_tables              | all                 | Collect from pg_statio_user_tables             |


### General Flags

| Name                | Description                                                              |
|---------------------|--------------------------------------------------------------------------|
| log.level           | Logging verbosity (default: info)                                        |
| web.listen-address  | Address to listen on for web interface and telemetry. (default: `:9135`) |
| web.telemetry-path  | Path under which to expose metrics. (default: `/metrics`)                |
| version             | Print the version information.                                           |
