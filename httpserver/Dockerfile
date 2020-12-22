# GRPC-demo
# build command : docker build . -t grpc-demo/server
# run command : docker run -it grpc-demo/server
FROM golang:latest

# Install grpc
RUN go get -u google.golang.org/grpc && \
    go get -u github.com/golang/protobuf/protoc-gen-go

# Install protoc and zip system library
RUN apt-get update && apt-get install -y zip && \
    mkdir /opt/protoc && cd /opt/protoc && wget https://github.com/protocolbuffers/protobuf/releases/download/v3.7.0/protoc-3.7.0-linux-x86_64.zip && \
    unzip protoc-3.7.0-linux-x86_64.zip

ENV PATH=$PATH:$GOPATH/bin:/opt/protoc/bin

# Copy the grpc proto file and generate the go module
RUN mkdir -p /go/src/protos
COPY ./protos/filepc.proto /go/src/protos
RUN cd /go/src && \
    protoc -I protos/ protos/filepc.proto --go-grpc_out=protos/ && \
    protoc -I=protos --go_out=protos protos/filepc.proto

# Copy the project to be executed
RUN mkdir -p /go/src/httpserver
COPY server.go /go/src/httpserver
EXPOSE 8081
EXPOSE 8082
ENTRYPOINT cd /go/src/httpserver && go run server.go 8081