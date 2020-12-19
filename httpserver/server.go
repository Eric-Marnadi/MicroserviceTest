package main

import (
	filepc "DSLabs/gowiki/protos"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net/http"
	"os"
	"strings"
)
var (
	serverAddr = "127.0.0.1:8082"
)
func parseRequest(r http.Request, prefix string) map[string]string {
	request := r.URL.Path[len(prefix):]
	tokens := strings.Split(request, "&")
	result := make(map[string]string)
	for _, token := range tokens {
		kv := strings.Split(token, "=")
		result[kv[0]] = kv[1]
	}
	return result
}
func viewHandler(w http.ResponseWriter, r *http.Request) {
	headers := parseRequest(*r, "/view/")
	filerequest := filepc.FileRequest{
		Filename:  headers["filename"],
		Extension: headers["extension"],
	}
	conn, _ := grpc.Dial(serverAddr)
	client := filepc.NewExplorerClient(conn)
	response, _ := client.GetFile(context.Background(), &filerequest)
	fmt.Fprintf(w, "<div>%s</div>", response.Content)
}
func main() {
	http.HandleFunc("/view", viewHandler)
	http.ListenAndServe(":" + os.Args[1], nil)
}
