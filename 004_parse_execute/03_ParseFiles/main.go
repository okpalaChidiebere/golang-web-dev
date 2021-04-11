package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	//We parsed one file
	tpl, err := template.ParseFiles("one.gmao")
	if err != nil {
		log.Fatalln(err)
	}

	//Execute that file to standard out
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	/*
		Parsed two more files

		Notice here that we called ParseFiles() again in the tpl variable that is a pointer to a template
		This method will add more files to the container that stores all the templates we parsed previously which is stored by the pointer tpl
		https://golang.org/pkg/text/template/#Template.ParseFiles
	*/
	tpl, err = tpl.ParseFiles("two.gmao", "vespa.gmao")
	if err != nil {
		log.Fatalln(err)
	}

	//ExecuteTemplate method allows you to execute a specific template when your pointer container tpl has more than one template files parsed in it
	err = tpl.ExecuteTemplate(os.Stdout, "vespa.gmao", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "two.gmao", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "one.gmao", nil)
	if err != nil {
		log.Fatalln(err)
	}

	//Execute executes the first file it finds in the container
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

}
