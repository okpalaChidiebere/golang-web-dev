package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/barred", barred)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	fmt.Print("Your request method at foo: ", req.Method, "\n\n")
}

func bar(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at bar:", req.Method)
	// process form submission here
	http.Redirect(w, req, "/", http.StatusTemporaryRedirect)
}

func barred(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at barred:", req.Method)
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

/*
In summary

TemporaryRedirect keeps the METHOD thesame. If a POST request comes in, the redirect will also be a POST

Request at foo '/' will be GET because its the home page
Request method for barried is GET (we make a request to load the form page)
Request method for bar will be POST (this does the form processiong and redirected with 303 which changes the method to GET)
Request method for foo will be POST (we redirected our server to this path)
*/
