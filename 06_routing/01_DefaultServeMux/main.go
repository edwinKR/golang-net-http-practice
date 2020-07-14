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

func p (res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "human~!!!!!")
}

func main() {
	// The Handle below requires a 'type Handler'. 'd' is a type Handler (and a type int).
	var d hotdog
	http.Handle("/dog", d)

	// The HandleFunc is attached to the DefaultServeMux.
	http.HandleFunc("/hello", h)
	http.HandleFunc("/world", w)

	// Handle vs HandleFunc
		// Handle: Takes a value of 'type Handler'
		// HandleFunc: Takes a func with the following signature ==> func(ResponseWriter, *Request)

	// http package Handle and HandleFunc are attached to the DefaultServeMux.

	// HandlerFunc vs HandleFunc
		// HandlerFunc is different from HandleFunc! Double check docs. We can do the below:
	http.Handle("/person", http.HandlerFunc(p))


	// When the nil is passed in as the handler, you should know immediately that this is using the DefaultServeMux.
	// If not nil, then somebody created their own a 'Multiplexer(Mux)'.
	http.ListenAndServe(":8080", nil)
}