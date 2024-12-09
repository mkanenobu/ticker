.PHONY: build
build:
	@go build -o bin/ticker cmd/ticker/main.go

.PHONY: run
run:
	@go run cmd/main.go

.PHONY: fmt
fmt:
	@go fmt ./...

.PHONY: install
install:
	go install ./...