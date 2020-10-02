#!/bin/bash

mkdir -p bin

REVISION=$(git rev-parse HEAD)
VERSION=$(git describe  --exact-match ${REVISION} 2>/dev/null)
BRANCH=$(git rev-parse --abbrev-ref HEAD)


set -x
CGO_ENABLED=0 go build \
    -mod=vendor -a -tags netgo \
    -ldflags "-X github.com/prometheus/common/version.Version=${VERSION}
        -X github.com/prometheus/common/version.Revision=${REVISION}
        -X github.com/prometheus/common/version.Branch=${BRANCH}
        -X github.com/prometheus/common/version.BuildUser=${USER}@$(hostname)
        -X github.com/prometheus/common/version.BuildDate=$(date +%Y%m%d-%X)" \
    -o bin/pg_exporter