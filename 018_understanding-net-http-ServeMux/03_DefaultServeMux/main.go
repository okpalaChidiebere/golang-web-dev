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
	var d hotdog
	var c hotcat

	/*
				here we are using the default Handler (DefaultServeMux type) instead of creating our own ServeMux using the http.NewServeMux()

				This is another way to set up your mux. But not the best way yet.



				Another way is to use the HandleFunc method at the http package level (https://golang.org/pkg/net/http/#HandleFunc)

				NOTE: func(ResponseWriter, *Request) you passed into the http.HandleFunc is different from HandlerFunc type (https://golang.org/pkg/net/http/#HandlerFunc)

				You have:
				func d ServeHTTP(res http.ResponseWriter, req *http.Request) {
			        io.WriteString(res, "cat cat cat")
		        }

				you pass the d function to the HandleFunc and not a Handler like:
				http.HandleFunc("/dog", d)

				Notice that d is not a Handler function. It is a function. Remember you have done this in typescript where you pass a function as an argument so that that function can be invoked in your function

				Using the HandleFunc is the best way to write a multiplexer for your server! Yeyyyy!

	*/
	http.Handle("/dog", d) //we use the Handle at the http package level (https://golang.org/pkg/net/http/#Handle) instead of Handle attached to a ServeMux
	http.Handle("/cat", c)

	http.ListenAndServe(":8080", nil) //You must pass in nil in the argument where you pass in your Handler variable type. "nil" means no value

}
