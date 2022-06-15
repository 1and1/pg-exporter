# Hacking on pg-exporter

## adding a new scraper

### define your model in collector/models

Models are defined by creating a struct like in the most ORM's. As pg-exporter uses [bun][1],
the same `bun` tags are used.

In addition, the following pg-exporter specific tags are available and required for code-generation:

- **help**: defines the help text for the metric
- **metric**: defines metric name and type. The first, optional paramter is the name. Type could be:
  - **counter**: for counter values
  - **gauge**: for gauge values
  - **label**: fo labels
  
### generate the full code for the model

run `go generate .` in the root folder of the project

### define the scraper

- create a new, empty, struct in the `collector` package, that implements the `Scraper` interface.
- Inside the `Scrape` function, fetch your data from the database with the given database connection.
  - see [bun][1] for details how to do that
- use the (generated) function `ToMetrics` on your result to provide the metrics to the prometheus handler


[1]: https://bun.uptrace.dev/