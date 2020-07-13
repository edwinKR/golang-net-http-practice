package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	//  Off of my req, parse the form. Remember! According to the docs, you can call 'req.Form()' or 'req.PostForm()' only after calling 'req.ParseForm()'
	err := req.ParseForm()

	if err != nil {
		// Remember, Fatalln includes the exit command.
		log.Fatalln(err)
	}
	fmt.Println(req.Form)
	tpl.ExecuteTemplate(w, "index.gohtml", req.Form)

}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var dog hotdog

	// Anything that comes in through the port 3030, will be handled by the Handler which is d in this case.
	http.ListenAndServe(":8080", dog)
}