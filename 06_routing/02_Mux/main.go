package main

import (
	"io"
	"net/http"
)

type dog int

func (d dog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "bark, bark bark")
}

type cat int

func (c cat) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "meow, meow, meow")
}

func main() {
	var d dog
	var c cat

	mux := http.NewServeMux()
	// Adding "/", what's the difference?
	// Try localhost:8080/dog/something/hello
	mux.Handle("/dog/", d)
	mux.Handle("/cat", c)

	http.ListenAndServe(":8080", mux)
}