package main

import (
	"io"
	"net/http"
)

/*
This code here is similar to the code in the /02 folder.

We did not use /assets ins /02 folder so you that you can fully understand what is going on
as we explained in the /02 folder!
*/
func main() {
	http.HandleFunc("/", dog)
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/assets/toby.jpg">`)
}
