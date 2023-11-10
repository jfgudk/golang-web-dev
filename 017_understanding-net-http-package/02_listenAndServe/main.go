package main

import (
	"fmt"
	"net/http"
)

type hotdog int

// this implements the http.Handler interface function ServeHTTP for the hotdog type
func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// this prints to the http.ResponseWriter which writes to the web client
	fmt.Fprintln(w, "Any code you want in this func")
	// prints request accept header
	fmt.Println(r.Header["Accept"])
}

func main() {
	var d hotdog
	// this serves the hotdog d that has a handler
	// use ListenAndServeTLS for HTTPS
	http.ListenAndServe(":8080", d)
}
