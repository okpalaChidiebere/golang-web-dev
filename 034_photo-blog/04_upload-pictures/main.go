package main

import (
	"crypto/sha1"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	uuid "github.com/satori/go.uuid"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	c := getCookie(w, req)
	// process form submission
	if req.Method == http.MethodPost {
		mf, fh, err := req.FormFile("nf")
		if err != nil {
			fmt.Println(err)
		}

		/*
			the File type returned by multipart is a bit different.
			https://golang.org/pkg/mime/multipart/#File

			But we should always call the close() function anytime it reading stuff
			https://golang.org/pkg/io/#Closer
		*/
		defer mf.Close()
		// create sha for file name
		ext := strings.Split(fh.Filename, ".")[1]          //we are getting the extension of the uploaded file and we get that from the filename. it could be .jpeg or .png
		h := sha1.New()                                    //create a new instance of the SHA
		io.Copy(h, mf)                                     //we copy the mf to the SHA instance
		fname := fmt.Sprintf("%x", h.Sum(nil)) + "." + ext //NOTE: calling the Sum() method will get you the actual sha of the file passed. %x converts the SHA to hexadecimal value then we concanate the extension to that hexadecimal value
		// create new file
		wd, err := os.Getwd() //we get the current working directory; basically the part on the server where this code is located
		if err != nil {
			fmt.Println(err)
		}
		path := filepath.Join(wd, "public", "pics", fname) //we end up having for eg "[some-directory-path-to-our-code]/public/pics/[some-sha-value].ext"
		nf, err := os.Create(path)                         //we create a file and give it the path we want it created on
		if err != nil {
			fmt.Println(err)
		}
		defer nf.Close()
		// copy
		mf.Seek(0, 0)   //we are reseting the ReadWrite head back to the begining of the file becuase when we did the sha hash, it read through the whole file(mf). So we have to reset it back to the begining of the file
		io.Copy(nf, mf) //copy the mf file the user uploaded to the new file that our OS created
		// add filename to this user's cookie
		c = appendValue(w, c, fname)
	}
	xs := strings.Split(c.Value, "|")          //split the cookie string value by the pipe character
	tpl.ExecuteTemplate(w, "index.gohtml", xs) //pass the slice of string to the template
}

func getCookie(w http.ResponseWriter, req *http.Request) *http.Cookie {
	c, err := req.Cookie("session")
	if err != nil {
		sID, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
	}
	return c
}

// takes in a file name now also
func appendValue(w http.ResponseWriter, c *http.Cookie, fname string) *http.Cookie {
	s := c.Value
	if !strings.Contains(s, fname) { //if the file name not in the cookie value we add it
		s += "|" + fname
	}
	c.Value = s
	http.SetCookie(w, c) //set the cookie
	return c             //return our cookie
}
