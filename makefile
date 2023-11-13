.PHONY: proto lint

proto:
	protoc -I ./api  --go_out ./api/serverpb --go_opt paths=source_relative --go-grpc_out ./api/serverpb --go-grpc_opt paths=source_relative ./api/server.proto

lint:
	golangci-lint run ./...