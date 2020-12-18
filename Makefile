.PHONY: protos
build:
	go build -o bin/server ingest/server.go
	go build -o bin/getfiles filesystem/getfiles.go
protos:
	protoc -I protos/ protos/filepc.proto --go-grpc_out=protos/
