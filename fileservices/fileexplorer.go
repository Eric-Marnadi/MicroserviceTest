package main

import (
	filepc "DSLabs/gowiki/protos"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"net"
	"os"
)
type FileExplorerServer struct {
}
func (s FileExplorerServer) GetFile(ctx context.Context, request *filepc.FileRequest) (*filepc.FileResponse, error) {
	body, _ := ioutil.ReadFile(request.Filename + "." + request.Extension)
	response := &filepc.FileResponse{
		Filename: request.Filename,
		Content:  body,
	}
	return response, nil
}
func (s *FileExplorerServer) mustEmbedUnimplementedExplorerServer() {}

func main() {
	address := "127.0.0.1:" + os.Args[1]
	listen, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	fmt.Printf("Server is listening on %v ...", address)
	s := grpc.NewServer()
	var fes filepc.ExplorerServer = FileExplorerServer{}
	filepc.RegisterExplorerServer(s, &fes)
	s.Serve(listen)
}
