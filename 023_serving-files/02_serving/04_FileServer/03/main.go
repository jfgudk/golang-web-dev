package main

import (
	"io"
	"net/http"
)

func main() {
	// this carves out URI / to this particular handler - which always returns toby.jpg
	http.HandleFunc("/", dog)
	// this carves out and causes anything in the /assets URI to be served from /assets folder, includes the directory listing
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/assets/toby.jpg">`)
}
