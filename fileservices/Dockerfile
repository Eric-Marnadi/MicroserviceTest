# GRPC-demo
# build command : docker build . -t grpc-demo/server
# run command : docker run -it grpc-demo/server
FROM golang:latest

# Install grpc
RUN go get -u google.golang.org/grpc && \
    go get -u github.com/golang/protobuf/protoc-gen-go && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

# Install protoc and zip system library
RUN apt-get update && apt-get install -y zip && \
    mkdir /opt/protoc && cd /opt/protoc && wget https://github.com/protocolbuffers/protobuf/releases/download/v3.7.0/protoc-3.7.0-linux-x86_64.zip && \
    unzip protoc-3.7.0-linux-x86_64.zip

ENV PATH=$PATH:$GOPATH/bin:/opt/protoc/bin

# Copy the grpc proto file and generate the go module
RUN mkdir -p /go/src/fileservices/protos
COPY filepc.proto /go/src/fileservices/protos
RUN cd /go/src/fileservices/
#RUN protoc -I=protos --go_out=protos protos/filepc.proto
#RUN protoc --go_out=. --go_opt=paths=source_relative \
#        --go-grpc_out=. --go-grpc_opt=paths=source_relative \
#        protos/filepc.proto
# Copy the project to be executed
COPY fileexplorer.go /go/src/fileservices
EXPOSE 8081
EXPOSE 8082
#ENTRYPOINT cd /go/src/fileservices && go run fileexplorer.go 8082
