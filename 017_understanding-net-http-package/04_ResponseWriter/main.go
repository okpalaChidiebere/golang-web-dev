package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	//We are basically setting any header we want
	w.Header().Set("Mcleod-Key", "this is from mcleod") //type Header has a lot of function attached to it. But we use Set() method to set headers that will be returned by the server https://golang.org/pkg/net/http/#Header
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, "<h1>Any code you want in this func</h1>")
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
