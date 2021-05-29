package main

import (
	"io"
	"net/http"
)

func main() {
	/*
		    Looking at the documentation https://golang.org/pkg/net/http/#FileServer, FileServer returns a handler that is why we are able to pass it into thehttp.Handle() method directly


			FileServer() method takles in a FileSystem as argument. We know FileSyetem is an interface looking at https://golang.org/pkg/net/http/#FileSystem
			We have a type Dir (https://golang.org/pkg/net/http/#Dir) which implements the FileSystem interface because it has the method "Open(name string) (File, error)" attached to it. So we are able to pass http.Dir() as argument to FileServer!

			The dot (.) we passed basically means the current directory where the main.go is.
			We are basically serving all the files in the /01 directory including the go files and image files as well

			Eg when we go to the endpoint (http://localhost:3000/), we will see all the files we are serving listed on the page
	*/
	http.Handle("/", http.FileServer(http.Dir("."))) //We are able to serve the image /toby.jpg because the FileServer has the files at the home route (/)
	http.HandleFunc("/dog/", dog)
	http.ListenAndServe(":8080", nil)
}

/*
When we go to /dog it, this function runs which asks for an image (/toby.jpg)
*/
func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/toby.jpg">`)
}
