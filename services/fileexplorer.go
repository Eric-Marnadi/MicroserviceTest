package main

import (
	filepc "DSLabs/gowiki/protos"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)
type FileExplorerServer struct {

}
func main() {
	address := "0.0.0.0:" + os.Args[1]
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	fmt.Printf("Server is listening on %v ...", address)

	s := grpc.NewServer()
	filepc.RegisterExplorerServer(s, &FileExplorerServer{})


	s.Serve(lis)
}
