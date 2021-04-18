package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/toby.jpg", dogPic)
	http.ListenAndServe(":8080", nil)
}

//At the root "/" we are asking for an image
func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/toby.jpg">`) //We are asking for the image my this tag calling for "/toby.jpg" as a result we run the code in the dogPic() method
}

func dogPic(w http.ResponseWriter, req *http.Request) {
	f, err := os.Open("toby.jpg") //we open the file as usual
	if err != nil {
		/* A lot of the time perople will use a constant in place for the third argument
		https://golang.org/pkg/net/http/#pkg-constants
		Eg isnatead of hard coding the status code our self like 404, we will use http.StatusNotFound
		*/
		http.Error(w, "file not found", 404)
		return
	}
	defer f.Close()

	fi, err := f.Stat() //This method returns a FileInfo and an error value https://golang.org/pkg/io/fs/#FileInfo
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}

	/*
		The FileInfo returned has different methods
		ServerContent method wants the ResponseWriter, Pointer to a Request, the filename, the lastime the file was modified and the actual file itself
		This method used the lastime the file as modified for the e-tag. The e-tag in HTTP is one of the ways cache is handled

		You will not often use this method for serving files in your server really but it is good to know!
		If you want to server your image based on cached time or modified time, this might be the function you want to use
	*/
	http.ServeContent(w, req, f.Name(), fi.ModTime(), f)
}
