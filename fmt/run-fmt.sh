#!/usr/bin/env bash
set -eu

docker build .. -t kyroy/gochecks
docker build . -t kyroy/gochecks-fmt

docker run --rm -v $PWD/../pkg/gotest/testdata/gotests/subpkgs:/t kyroy/gochecks-fmt --dir /t
