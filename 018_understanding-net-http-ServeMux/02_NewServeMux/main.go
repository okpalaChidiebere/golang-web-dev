package main

import (
	"io"
	"net/http"
)

type hotdog int

func (d hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog dog")
}

type hotcat int

func (c hotcat) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat cat cat")
}

func main() {
	//We basically created two different Handler and passed then the mux
	var d hotdog
	var c hotcat

	mux := http.NewServeMux() //This returns a pointer to ServerMux type(*ServeMux)
	/* This way of writing mux is not totally elegant but a bit better than 01_routing */
	mux.Handle("/dog/", d) //what mux does is based on some certain method and path or route, we will run a certain chink of code. Old School way like PHP you run a file instead of code.
	mux.Handle("/cat", c)

	/*
		If a request is made at /dog/somehing or /dog we will run the dog code

		if a request comes in at /cat/something/else that cat code will not run because there is no downline. So the client will get a 404
	*/

	http.ListenAndServe(":8080", mux)
}
