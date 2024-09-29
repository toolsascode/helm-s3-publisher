#!/bin/sh
set -e
rm -rf manpages
mkdir manpages
go run ./cmd/helm-s3-publisher man | gzip -c -9 > manpages/helm-s3-publisher.1.gz