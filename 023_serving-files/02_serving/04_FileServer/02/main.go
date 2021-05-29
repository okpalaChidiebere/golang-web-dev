package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", dog)
	/*
		Any request coming in at the /resources path will use the Handler (http.StripPerfix method)

		http.StripPrefix method strips off "/resources" from the request path "/resources/toby.jpg". So we now have "/toby.jpg"

		The httpFileServer will now look for the image "/toby.jpg". You can image the Dir method looking for your image like http.Dir("./assets/toby.jpg")
	*/
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/resources/toby.jpg">`) // /resources/toby.jpg is one of the image served at the path /resources
}

/*

./assets/toby.jpg

Your code main.go is outside the assets folder, so it will not get served
*/
