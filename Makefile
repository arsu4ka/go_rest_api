.PHONY: build
build:
	go build -o bin/apiserver ./cmd/apiserver

.PHONY: run
run:
	go run ./cmd/apiserver

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := build