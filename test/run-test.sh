#!/usr/bin/env bash
set -eu

docker build .. -t kyroy/gochecks
docker build . -t kyroy/gochecks-test

docker run --rm -v $PWD/../pkg/gotest/testdata/gotests/test1:/t kyroy/gochecks-test --dir=/t
