package main

import (
	"fmt"
	//"github.com/satori/go.uuid"
	"net/http"
)

// For this code to run, you will need this package:
// go get github.com/satori/go.uuid

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("session")
	if err != nil {
		// If there is an error, we have no cookie; so we create a new one

		id := uuid.NewV4() //ideally, the package name is thesame as the folder name, but satori went agains the rule :) The foldername is "go.uuid" and package name is "uuid"
		cookie = &http.Cookie{
			Name:  "session",
			Value: id.String(), //we take the uuid and turn it to a string
			// Secure: true, //if we were using HTTPS, we will uncomment this line. This allows the cookie to be sent only on https securely
			HttpOnly: true, //HTTPOnly means that you cant access this cookie with JavaScript. This makes this cookie even more secure
			Path:     "/",  //we set th path to be default which allows the cookie to be accessed all over our website. But we can specify a special subpath as well eh /auth/blaBla
		}
		http.SetCookie(w, cookie)
	}
	fmt.Println(cookie) //we print the cookie on log server if there is one!
}
