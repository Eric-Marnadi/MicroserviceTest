package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type Page struct {
	Title string
	Body  []byte
}
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.	ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}
func parseRequest(r http.Request, prefix string) (map[string]string) {
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
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<div>%s</div>", p.Body)
}
func createHandler(w http.ResponseWriter, r *http.Request) {
	m := parseRequest(*r, "/create/")
	fmt.Fprintf(w, "<div>%s</div>", m)
}
func main() {
	http.HandleFunc("/create/", createHandler)
	http.HandleFunc("/view/", viewHandler)
	log.Fatal(http.ListenAndServe(":" + os.Args[1], nil))
}
