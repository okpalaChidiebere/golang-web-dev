package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	v := req.FormValue("q") //We still get the form value for q this time but it is send through the request body and not url all because we changed the method to POST
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
	<form method="POST">
	 <input type="text" name="q">
	 <input type="submit">
	</form>
	<br>`+v) //we print the form data submitted right below the form.
}
