package main

import (
	"html/template"
	"log"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	/*
		From the documentation https://golang.org/pkg/net/http/#Request, we know the Request type is a struct and has methods attached to it

		Now, to get the Form or PostForm field in the struct, it is said you have to call the ParseForm() method

		The ParseFrom() method is thod that you can call on a variable that is of type Pointer to a Request(*Request)
		https://golang.org/pkg/net/http/#ParseForm
	*/
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	/*
		NOTICE here how we access the "Form" that is in the req struct!
		When you click on the Form value in the documentation showing the Request struct, you will see that the Form value is a of type map of string  https://golang.org/pkg/net/url/#Values

		Remember in the template file when it is submited we passed a queryString in the url and the form data in the body of the request, by getting the Form filed in the request, we will have fname variable be a slice that has ["james", "whateverYourInputIs"]
		Form filed will get data from the url passed as queryString and payload for the request from the form body
		Input tags names in html are always submitted as slice of strings anyways where there is just one element is the slice or not

		We pass that fileld as data to our template because we want to print the form input back to the user. We HAVE to range over the From value passed as data in our template

		Ideally, we may perform a database operation or something
	*/
	tpl.ExecuteTemplate(w, "index.gohtml", req.Form) //we could have called req.PostForm but PostForm only gets formdata in the body of the request but we want data passed in the url as well
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d) //When any thing comes in at port8080, it will be handled by the ServeHTTP() method
}
