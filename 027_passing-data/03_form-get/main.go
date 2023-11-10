package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	// this reads the URI query paramater `q`
	v := req.FormValue("q")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// form field using GET returns entered value in URI as a param value as well as the returned body
	// this creates a feedback from the form value to the param value
	io.WriteString(w, `
	<form method="GET">
	 <input type="text" name="q">
	 <input type="submit">
	</form>
	<br>`+v)
}
