.PHONY: all build test run-example

all: build test run-example

build:
	go build -o tsgen ./cmd/tsgen

test:
	go test ./...

run-example: build
	./tsgen -dsl ./examples/cli/dsl/fileio.dsl
	go test ./examples/cli/tests
