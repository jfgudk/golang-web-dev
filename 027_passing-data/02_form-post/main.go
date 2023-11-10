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
	// POST form values take precendence over query param values
	v := req.FormValue("q")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// form field for q value with submit button
	// displays form value in body in the return of the POST command only
	io.WriteString(w, `
	<form method="POST">
	 <input type="text" name="q">
	 <input type="submit">
	</form>
	<br>`+v)
}
