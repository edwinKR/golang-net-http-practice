package main

import (
	"io"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "bark bark bark!")
}

func h (res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "hello~~~~")
}

func w (res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "world!!!!")
}

func main() {
	// The Handle below requires a 'type Handler'. 'd' is a type Handler (and a type int).
	var d hotdog
	http.Handle("/dog", d)

	// The HandleFunc is attached to the DefaultServeMux.
	http.HandleFunc("/hello", h)
	http.HandleFunc("/world", w)

	// Handle vs HandleFunc ==> HandleFunc does NOT require the Handler type! It requires a handler function in a certain format.
	// http package Handle and HandleFunc are attached to the DefaultServeMux.

	// When the nil is passed in as the handler, you should know immediately that this is using the DefaultServeMux.
	// If not nil, then somebody created their own a 'Multiplexer(Mux)'.
	http.ListenAndServe(":8080", nil)
}