GOCMD = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOTEST = $(GOCMD) test
DEP = dep
GIT = git

BINARY = "bin/sida"
VERSION = `$(GIT) describe --tags`
TIME = `date +%FT%T%z`

all: test build
clean:
	$(GOCLEAN)
	rm -f $(BINARY)
build:
	$(GOBUILD) -o $(BINARY) -ldflags "-X main.Version=${VERSION} -X main.Build=${TIME}" -v cmd/sida/*.go
test:
	$(GOTEST) ./...
cover:
	$(GOTEST) --cover ./...
deps:
	$(DEP) ensure
