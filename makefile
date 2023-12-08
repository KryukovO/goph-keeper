.PHONY: mockgen build proto lint

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

mockgen:
	mockgen -destination internal/server/repository/mocks/user.go -package mocks github.com/KryukovO/goph-keeper/internal/server/repository UserRepository
	mockgen -destination internal/server/repository/mocks/auth.go -package mocks github.com/KryukovO/goph-keeper/internal/server/repository AuthDataRepository
	mockgen -destination internal/server/repository/mocks/text.go -package mocks github.com/KryukovO/goph-keeper/internal/server/repository TextDataRepository
	mockgen -destination internal/server/repository/mocks/bank.go -package mocks github.com/KryukovO/goph-keeper/internal/server/repository BankDataRepository
	mockgen -destination internal/server/repository/mocks/subscription.go -package mocks github.com/KryukovO/goph-keeper/internal/server/repository SubscriptionRepository
	mockgen -destination internal/server/filestorage/mocks/storage.go -package mocks github.com/KryukovO/goph-keeper/internal/server/filestorage FileStorage

test:
	go test -v -timeout 30s -race ./...

cover:
	go test -timeout 30s -race -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out
	rm coverage.out

cover-html:
	go test -timeout 30s -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	rm coverage.out
