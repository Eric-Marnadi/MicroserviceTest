build:
	go build -o bin/server server.go
	go build -o bin/getfiles getfiles.go
run:
	go run server.go
protos:
	protoc -I protos/ protos/filepc.proto --go-grpc_out=protos/