.PHONY: protos
build:
	go build -o httpserver/server httpserver/server.go
	go build -o fileservices/getfiles fileservices/fileexplorer.go
protos:
	protoc -I protos/ protos/filepc.proto --go-grpc_out=protos/
	protoc -I=protos --go_out=protos protos/filepc.proto
