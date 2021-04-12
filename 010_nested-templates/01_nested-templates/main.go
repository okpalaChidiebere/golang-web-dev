package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

/*
Nothing really changes in the go file from what you dont already know
like Parsing your template files, Executing them and passing data to them

The only thinng new now is that some of your template file are now being used inside other files!

Some template file just contain logic of you defining them!
*/

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", 42)
	if err != nil {
		log.Fatalln(err)
	}
}
