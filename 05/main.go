package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Some-Key", "Some Value says hello!")
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	//w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, "<h3>Any code you want in this func</h3>")
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}