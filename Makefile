.PHONY: protos
build:
	go build -o bin/server httpserver/server.go
	go build -o bin/getfiles services/fileexplorer.go
protos:
	protoc -I protos/ protos/filepc.proto --go-grpc_out=protos/
	protoc -I protos/ protos/hellopb.proto --go-grpc_out=protos/
