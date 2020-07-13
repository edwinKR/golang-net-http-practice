package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
	"fmt"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Println(err)
	}

	data := struct {
		Method			string
		URL				*url.URL
		Submissions		url.Values
		Header			http.Header
		ContentLength	int64
	}{
		req.Method,
		req.URL,
		req.Form,
		req.Header,
		req.ContentLength,
	}

	fmt.Printf("%+v\n", data)
	tpl.ExecuteTemplate(w, "index.gohtml", data)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}