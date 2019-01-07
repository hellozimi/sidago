DEP=dep
GIT=git

PROJECTNAME=$(shell basename "$(PWD)")
BINARY=bin/sida
VERSION=`$(GIT) describe --tags`
TIME=`date +%FT%T%z`

MAKEFLAGS += --silent
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${TIME}"

.PHONY: help
all: help
help: Makefile
	@echo " Choose a command to run in "$(PROJECTNAME)":"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'

## clean: Cleans the project and build directory
clean:
	@go clean
	@rm -f $(BINARY)
## build: Builds a binary
build:
	@go build -o $(BINARY) $(LDFLAGS) -v cmd/sida/*.go
## test: Run all tests
test:
	@go test ./...
## cover: Run all tests and report coverage
cover:
	@go test --cover ./...
## deps: Installs dependencies
deps:
	@dep ensure
## install: Installs the binary
install: 
	@go build -i -o $(BINARY) $(LDFLAGS) -v cmd/sida/*.go