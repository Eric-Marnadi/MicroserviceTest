package main

import (
	"net/http"
	"os"
)
func viewHandler(w http.ResponseWriter, r *http.Request) {

}
func main() {
	http.Handle("/view", viewHandler)
	http.ListenAndServe(":" + os.Args[1], nil)
}
