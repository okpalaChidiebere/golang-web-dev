package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	/*
		In practise, you use ExecuteTemplate() over Execute() and
		you use ParseGlob() over ParseFiles()
		https://golang.org/pkg/text/template/
	*/

	//This will parse all the templates files in templates folder
	tpl, err := template.ParseGlob("templates/*")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil)
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
