package main

import (
	"html/template" //NOTE here we now use html package instead of Text template. All the functions that exists in text template still carries over here. Nothing new to learn. But this one works better if you are actually building real web pages for your site over the text/template package!
	"log"
	"os"
)

type Page struct {
	Title   string
	Heading string
	Input   string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	home := Page{
		Title:   "Escaped",
		Heading: "Danger is escaped with html/template",
		Input:   `<script>alert("Yow!");</script>`, //This will not execute
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", home)
	if err != nil {
		log.Fatalln(err)
	}

	/*
		https://pkg.go.dev/html/template has more additional functionalities for doing
		html pages. It has a lot of safety for preventing web attacks like Cross-Site Scripting.
		So it escapes a lot of characters
	*/
}
