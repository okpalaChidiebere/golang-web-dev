package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/toby.jpg", dogPic) //We have a route for the request to the image that will run a certain code
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w, `
	<img src="/toby.jpg">
	`) //We are asking for an image that the source is toby. We are asking for toby at the root, We are more absolute thats why you have the slash "/" before the image name. "toby.jpg" is equal to "/toby.jpg"
}

func dogPic(w http.ResponseWriter, req *http.Request) {
	f, err := os.Open("toby.jpg") //We open that file. Pointer to a file(*FIle) and error values are returned
	if err != nil {               //We deal with our error, if we have one
		http.Error(w, "file not found", 404)
		return
	}
	defer f.Close()

	io.Copy(w, f) //write the File to our ResponseWriter using io.Copy
}
