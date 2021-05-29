package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

/*
We get open the file that was submitted by the user, read it and then write the content
of the file back to the response
*/
func foo(w http.ResponseWriter, req *http.Request) {

	var s string
	fmt.Println(req.Method)
	//we are checking if the request method is POST. NOTE that http.MethodPost is a constant
	//https://golang.org/pkg/net/http/#pkg-constants
	if req.Method == http.MethodPost {

		// open
		f, h, err := req.FormFile("q") //FromFile method is used get form input file values! . The identifier for this file value is "q" Remember how we use FormValue method to get Form input text values
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()
		//Side NOTE: you can Reac, Seek and Close this file. If you how to seek the file, you have to use te Seeker type

		// for your information
		fmt.Println("\nfile:", f, "\nheader:", h, "\nerr", err)

		// read
		bs, err := ioutil.ReadAll(f) //file has the reader interface, this is why we can pass it to ReadAll method
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		s = string(bs)
	}

	/*
		This block of code will run whether it is a POST or GET request coming in

		NOTE: we did not specify an action url for this form we are displaing.
		This means it will post the action to this page where the request is coming from /
		We dont have to specify the action attribute if we want the form to be subitted at thesame page!
	*/
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
	<form method="POST" enctype="multipart/form-data">
	<input type="file" name="q">
	<input type="submit">
	</form>
	<br>`+s) //The "s" variable we write back as request will either contain a zero value or content of the file the user uploaded
}
