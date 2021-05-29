package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	//We parse our file and get our pointer to a template
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	//create a new file
	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatalln("error creating file", err)
	}
	defer nf.Close()

	//Dump the parsed file results into the new file we created
	//NOTE: we are not writing the results to an standardOutput(terminal) this time. We are writing to a file. nf is of File type and File type implements the Writer interface
	err = tpl.Execute(nf, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

// Do not use the above code in production
// We will learn about efficiency improvements soon!
