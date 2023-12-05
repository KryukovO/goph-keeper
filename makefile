.PHONY: build proto lint

BUILDDATE=$(shell date +'%d-%m-%Y')
BUILDVERSION=v0.0.1

proto:
	protoc -I ./api  --go_out ./api/serverpb --go_opt paths=source_relative --go-grpc_out ./api/serverpb --go-grpc_opt paths=source_relative ./api/server.proto

lint:
	golangci-lint run ./...

build:
	# Linux
	GOOS=linux GOARCH=amd64 go build -o build/client/client_linux_amd64 -ldflags "-X main.buildVersion=${BUILDVERSION} -X main.buildDate=${BUILDDATE}" cmd/client/main.go
	GOOS=linux GOARCH=amd64 go build -o build/server/server_linux_amd64 -ldflags "-X main.buildVersion=${BUILDVERSION} -X main.buildDate=${BUILDDATE}" cmd/server/main.go
	# Windows
	GOOS=windows GOARCH=amd64 go build -o build/client/client_win_amd64.exe -ldflags "-X main.buildVersion=${BUILDVERSION} -X main.buildDate=${BUILDDATE}" cmd/client/main.go
	GOOS=windows GOARCH=amd64 go build -o build/server/server_win_amd64.exe -ldflags "-X main.buildVersion=${BUILDVERSION} -X main.buildDate=${BUILDDATE}" cmd/server/main.go
	# MacOS
	GOOS=darwin GOARCH=amd64 go build -o build/client/client_darwin_amd64 -ldflags "-X main.buildVersion=${BUILDVERSION} -X main.buildDate=${BUILDDATE}" cmd/client/main.go
	GOOS=darwin GOARCH=amd64 go build -o build/server/server_darwin_amd64 -ldflags "-X main.buildVersion=${BUILDVERSION} -X main.buildDate=${BUILDDATE}" cmd/server/main.go