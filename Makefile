.DEFAULT_GOAL := run
fmt:
	go fmt ./...
.PHONY:fmt
lint: fmt
	golangci-lint run
.PHONY:lint
vet: fmt
	go vet ./...
	shadow ./...
.PHONY:vet
run: vet
	go run .
.PHONY:run
build: vet
	go build
.PHONY:build
