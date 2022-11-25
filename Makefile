.DEFAULT_GOAL := run
fmt:
	go fmt ./...
.PHONY:fmt
lint: fmt
	golint ./...
	golangci-lint run
.PHONY:lint
vet: fmt
	go vet ./...
	shadow ./...
.PHONY:vet
run: vet
	go run main.go
.PHONY:run
