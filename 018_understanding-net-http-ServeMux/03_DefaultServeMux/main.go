package main

import (
	"io"
	"net/http"
)

type hotdog int

func (d hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog dog")
}

type hotcat int

func (c hotcat) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat cat cat")
}

func main() {
	var d hotdog
	var c hotcat

	// this is a handle to the DefaultServeMux
	http.Handle("/dog", d)
	http.Handle("/cat", c)

	// nil tells serve to use DefaultServerMux
	http.ListenAndServe(":8080", nil)
}
