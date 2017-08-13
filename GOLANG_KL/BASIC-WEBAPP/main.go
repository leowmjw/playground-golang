package main

import (
	"net/http"
	"fmt"
	"html"
)

type httpserver struct {
	id int
}

func indexHandler(w http.ResponseWriter, req *http.Request) (http.Handler) {
	return nil
}

func messageHandler(w http.ResponseWriter, req *http.Request) {

}

func newmessageHandler(w http.ResponseWriter, req *http.Request) {

}

func (this *httpserver) ServeHTTP(w http.ResponseWriter, req *http.Request) {

}

func main() {

	http.Handle("/foo", httpserver{})

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	http.ListenAndServe(":9090", nil)
}
