server:
	@go run cmd/main/main.go

client:
	@go run cmd/client/client.go

proto:
	@protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=module=github.com/vitorsalgado/grpc-go cmd/main/proto/*.proto
