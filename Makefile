GOCMD = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOTEST = $(GOCMD) test
DEP = dep


BINARY = "bin/sida"

all: test build
clean:
	$(GOCLEAN)
	rm -f $(BINARY)
build:
	$(GOBUILD) -o $(BINARY) -v cmd/sida/*.go
test:
	$(GOTEST) ./...
cover:
	$(GOTEST) --cover ./...
deps:
	$(DEP) ensure
