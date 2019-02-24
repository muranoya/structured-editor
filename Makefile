NAME=se

GOCMD=go
GOBUILD=$(GOCMD) build
GOFMT=gofmt
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test

BINARY=$(NAME)

.PHONY: all test build clean check fmt
all: fmt check test build

fmt:
	$(GOFMT) -s -l -e -w .

check:
	errcheck -exclude errcheck_excludes.txt -asserts -verbose ./...
	go vet ./...
	golint src/...

test:
	$(GOTEST) -v ./...

build:
	$(GOBUILD) -o bin/$(BINARY) src/main.go

clean:
	$(GOCLEAN)
	rm -fr bin/
