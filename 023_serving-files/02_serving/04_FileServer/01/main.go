package main

import (
	"io"
	"net/http"
)

func main() {
	// this serves all files in this dir
	http.Handle("/", http.FileServer(http.Dir(".")))
	// this carves out URI /dog to this particular handler
	http.HandleFunc("/dog/", dog)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/toby.jpg">`)
}
