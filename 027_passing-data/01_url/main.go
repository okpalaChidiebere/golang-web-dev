package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	v := req.FormValue("q")             //FormValue method is used to get each variable passed from the url. In this case "q" is the variable which has the value "dog"
	fmt.Fprintln(w, "Do my search: "+v) //we will this string in the web browser "Do my search: dog"
}

// visit this page:
// http://localhost:8080/?q=dog
