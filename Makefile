generate:
	protoc -I . ./api/*.proto -I ./api --go-grpc_out=internal/pb --go_out=internal/pb