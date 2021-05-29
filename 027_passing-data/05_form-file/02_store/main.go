package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

/*
This time we open the file uploaded, read the file and write the
contents of the uploaded file to another file and save it in our server
*/
func foo(w http.ResponseWriter, req *http.Request) {

	var s string
	if req.Method == http.MethodPost {

		// open
		f, h, err := req.FormFile("q")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()

		// for your information
		fmt.Println("\nfile:", f, "\nheader:", h, "\nerr", err)

		// read
		bs, err := ioutil.ReadAll(f)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		s = string(bs) //we convert our byte slice to a string. The return the content of the file uploaded in the request and print in the webpage

		// store on server
		dst, err := os.Create(filepath.Join("./user/", h.Filename)) //we created a file in of name ./user/<NAME_OF_THE_FILE_USER_UPLOADED>
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		//ds varible is a pointer to the file we created, so we can write directly to that file! We call the Write() method on that variable
		_, err = dst.Write(bs) //We write the byte slice(content of the uploaded file) to the new file we created
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tpl.ExecuteTemplate(w, "index.gohtml", s)
}
