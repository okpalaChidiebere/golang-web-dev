package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

/*
init() initializes your program. It runs only once.

So for more performance of our code, we want to parse our template files only once

Our files will be parsed when our program is starting up
*/
func init() {
	//The Must() method does error checking for us. We dont have to write the error checking code ourself! It takes a pointer to a Template and and an error value which the ParseGlob() method returns
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	err := tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "vespa.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
