package main

import (
	"io"
	"net/http"
)

func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog dog")
}

func c(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat cat cat")
}

func main() {

	/*
		Handle wants a route and a Handler. We know HandlerFunc implements the Handler interface (which the Handle() method wants) because it has the ServerHTTP method attached to it; but d is just a function but it matches the underlying type for a HandlerFunc type.
		So we use type conversion to convert the d which is just a function to HandlerFunc.

		https://golang.org/pkg/net/http/#HandlerFunc
	*/
	http.Handle("/dog", http.HandlerFunc(d))
	http.Handle("/cat", http.HandlerFunc(c))

	http.ListenAndServe(":8080", nil)
}

// this is similar to this:
// https://play.golang.org/p/X2dlgVSIrd
// ---and this---
// https://play.golang.org/p/YaUYR63b7L
