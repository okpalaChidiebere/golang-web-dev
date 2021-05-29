package main

import (
	"io"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	//This is one way to create a multiplexer. This works but there is a better way to write this code
	//This is the least Elegent way to write your multiplexer
	switch req.URL.Path {
	case "/dog":
		io.WriteString(w, "doggy doggy doggy")
	case "/cat":
		io.WriteString(w, "kitty kitty kitty")
	}
}

func main() {
	var d hotdog //NOTE that d variable is set to zero value here. d is of type Handler, hotdog and underline its an int :)
	http.ListenAndServe(":8080", d)
}
