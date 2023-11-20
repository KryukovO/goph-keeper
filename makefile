.PHONY: build proto lint

BUILDDATE=$(shell date +'%d-%m-%Y')
BUILDVERSION=v0.0.1

proto:
	protoc -I ./api  --go_out ./api/serverpb --go_opt paths=source_relative --go-grpc_out ./api/serverpb --go-grpc_opt paths=source_relative ./api/server.proto

lint:
	golangci-lint run ./...

build:
	go build -o cmd/client/client -ldflags "-X main.buildVersion=${BUILDVERSION} -X main.buildDate=${BUILDDATE}" cmd/client/main.go
	go build -o cmd/server/server -ldflags "-X main.buildVersion=${BUILDVERSION} -X main.buildDate=${BUILDDATE}" cmd/server/main.go