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
	fmt.Print("Your request method at foo: ", req.Method, "\n\n") //we will confirm that GET method was send here because of the redirect!
}

/*
We could proccess the from sent from the '/barred' path, but NO!
we want to redirect the request to '/'. The '/' runs the foo() method
*/
func bar(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at bar:", req.Method)
	// process form data
	w.Header().Set("Location", "/")    //NOTICE that here we are explicitly setting the "Location" which is going to be "/" in the response header
	w.WriteHeader(http.StatusSeeOther) //SeeOther means the method is always going to be GET and it going to send it to the location "/"
}

/*
When we visit the /bar, it xecuts a tempplate that has a form on it, and when we submit that form to /bar
*/
func barred(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at barred:", req.Method)
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

/*
In summary

The 303 always sends a GET request to the new path to redirect to! it does
not matter if the matter is the request that comes in initially was a POST
or GET request

Request method for barried is GET (we make a request to load the form page)
Request method for bar will be POST (this does the form processiong and redirected with 303 which changes the method to GET)
Request method for foo will be GET (we redirected our server to this path)
*/
