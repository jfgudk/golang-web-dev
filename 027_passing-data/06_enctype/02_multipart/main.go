package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

type person struct {
	FirstName  string
	LastName   string
	Subscribed bool
}

// code in example 01, 02 and 03 is identical
// templates are identical except this is using
// <form method="POST" enctype="multipart/form-data">
// enctype="multipart/form-data" is need for file uploading
// input results in output in form:
// ------WebKitFormBoundaryS8CEzoFKvJrESetc Content-Disposition: form-data; name="first" Jer ------WebKitFormBoundaryS8CEzoFKvJrESetc Content-Disposition: form-data; name="last" Gud ------WebKitFormBoundaryS8CEzoFKvJrESetc--
func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {

	// body
	bs := make([]byte, req.ContentLength)
	req.Body.Read(bs)
	body := string(bs)

	err := tpl.ExecuteTemplate(w, "index.gohtml", body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatalln(err)
	}
}
