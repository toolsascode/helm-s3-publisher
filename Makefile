# Makefile

GOBASE := $(shell pwd)
GOBIN := $(GOBASE)/bin

.PHONY: help
all: help

# .PHONY: dev
# dev:
# 	docker compose -f ./example/compose.yaml up --remove-orphans

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: dist-test
dist-test:
	export GPG_FINGERPRINT=$(gpg --list-secret-keys --keyid-format=long | grep --only-matching --extended-regexp "[[:xdigit:]]{40}")
	goreleaser --snapshot --clean --skip=publish

.PHONY: test
test:
	go test -coverprofile=coverage.out -cover ./... && go tool cover -html=coverage.out -o coverage.html

.PHONY: docs
docs:
	go run $(GOBASE)/cmd/helm-s3-publisher docs

.PHONY: run
run:
	go run $(GOBASE)/cmd/helm-s3-publisher argo /Users/carlosjunior/projects/helm-charts --force --git-ls-tree --exclude-paths ".git, .github" --log-level debug


.PHONY: build
build:
	golangci-lint run
	go build -v -ldflags="-X 'main.Version=v1.0.0-beta' -X 'main.commit=$(shell git rev-parse --short HEAD)' -X 'main.builtBy=$(shell id -u -n)' -X 'main.date=$(shell date)'" $(GOBASE)/cmd/helm-s3-publisher

.PHONY: version
version:
	go run -ldflags="-X 'main.version=v1.0.0-beta' -X 'main.commit=$(shell git rev-parse --short HEAD)' -X 'main.builtBy=$(shell id -u -n)' -X 'main.date=$(shell date)'" $(GOBASE)/cmd/helm-s3-publisher version
	@echo
	go run -ldflags="-X 'main.version=v1.0.0-beta' -X 'main.commit=$(shell git rev-parse --short HEAD)' -X 'main.builtBy=$(shell id -u -n)' -X 'main.date=$(shell date)'" $(GOBASE)/cmd/helm-s3-publisher -v
.PHONY: help
help: Makefile
	@echo
	@echo "Usage: make [options]"
	@echo
	@echo "Options:"
	@echo "    build     Create binary file"
	@echo "    run       Run helm-s3-publisher"
	@echo "    dist-test Run for Releaser test"
	@echo "    docs  	 Generete docs"
	@echo "    version   Set version in go application"
	@echo "    Help	"
