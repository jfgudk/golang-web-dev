package main

import (
	"html/template"
	"log"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// parse form required before request values are fully available
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	// submitted form content is accessed via req.Form (body and URL params) and/or PostForm (body only)
	// alse see MultipartForm
	tpl.ExecuteTemplate(w, "index.gohtml", req.Form)
	// Form is a url.Values type which is a map of strings
	// see http.Request for request struct
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
