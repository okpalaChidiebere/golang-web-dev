package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	/*
		The ParseFiles() method returns a pointer to a template and error values

		Think of pointer to a template as a container that holds all the files you have parsed. In this case its is just one
		In the ParseFiles method you can parse mor than one file eg ParseFiles("tpl.gohtml", "tpl2.gohtml", ...)
	*/
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	//since the tp, is a pointer to a template we can now call the Execute function https://golang.org/pkg/text/template/
	//in the Execute method, we care passing a Wirter interface(os.StdOut which is our terminal which we choose) and no data (nil mease empty value)
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

// Do not use the above code in production
// We will learn about efficiency improvements soon!
