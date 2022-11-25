.DEFAULT_GOAL := run
fmt:
	go fmt ./...
.PHONY:fmt
lint: fmt
	golangci-lint run
.PHONY:lint
run: lint
	go run .
.PHONY:run
build: lint
	go build
.PHONY:build
