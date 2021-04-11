package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	/*
		Rememeber in the previous programs before this point, we use to pass nil
		in the third argument which is for data. Now we pass in 42. The 42 will be populated in the
		"{{.}}" of the template file(tpl.gohtml).

		You only get to pass in one pice of data though. The data can be of type slice, map or another struct
	*/
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", 42)
	if err != nil {
		log.Fatalln(err)
	}
}
