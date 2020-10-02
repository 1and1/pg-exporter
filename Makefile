REVISION   ?= $(shell git rev-parse HEAD)
BRANCH     ?= $(shell git rev-parse --abbrev-ref HEAD)
VERSION    ?= $(shell git describe  --exact-match $(REVISION) 2>/dev/null)
BUILD_USER ?= $(USER)@$(shell hostname)
BUILD_DATE ?= $(shell date +%Y%m%d-%T)
all: bin/pg-exporter

bin/pg-exporter:
	mkdir -p bin
	CGO_ENABLED=0 go build \
		-mod=vendor -a -tags netgo \
		-ldflags "-X github.com/prometheus/common/version.Version=$(VERSION) \
			-X github.com/prometheus/common/version.Revision=$(REVISION) \
			-X github.com/prometheus/common/version.Branch=$(BRANCH) \
			-X github.com/prometheus/common/version.BuildUser=$(BUILD_USER) \
			-X github.com/prometheus/common/version.BuildDate=$(BUILD_DATE)" \
			-o bin/pg_exporter
