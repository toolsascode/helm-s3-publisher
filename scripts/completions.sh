#!/bin/sh
set -e
rm -rf completions
mkdir completions
for sh in bash zsh fish; do
	go run ./cmd/helm-s3-publisher completion "$sh" > "completions/helm-s3-publisher.$sh"
done