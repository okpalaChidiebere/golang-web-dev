package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	data := struct {
		/*
			We define our own type and the fields that we are getting from the req
			struct that is returned

			No magic here the filed types have to be of thesame type as returned by the Request
			Eg *url.URL in our struct is also thesame in the req struct*/
		Method      string
		URL         *url.URL
		Submissions url.Values
	}{
		//We then initialze these fields
		req.Method,
		req.URL, //We get the URL or path field from the req struct returned
		req.Form,
	}
	tpl.ExecuteTemplate(w, "index.gohtml", data)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
