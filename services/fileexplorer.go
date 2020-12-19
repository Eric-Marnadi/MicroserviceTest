package main

import (
	. "DSLabs/gowiki/protos"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)
// HelloServiceServer is the server API for HelloService service.
type FileExplorerServer interface {
	Hello(context.Context, *FileRequest) (*FileResponse)
}

func main() {
	address := "127.0.0.1:" + os.Args[1]
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	fmt.Printf("Server is listening on %v ...", address)

	s := grpc.NewServer()
	filepc.RegisterExplorerServer(s, &FileExplorerServer{})
	s.Serve(lis)
}
