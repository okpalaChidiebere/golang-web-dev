package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", foo)
	/*
		FYI: we are handling favicon request because when your browser asks for
		this icon, and its fails, it will go to the default route '/' of
		which will increament our counter and it is a wrong count

		We want to count every other request apart from the favicon.ico request
	*/
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {

	cookie, err := req.Cookie("my-cookie")

	/*
		http.ErrNoCookie is an error returned when a cookie is not found

		So if there is no cookie set by our site already, we set one.
	*/
	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "my-cookie",
			Value: "0",
			Path:  "/",
		}
	}

	//We get the value of the cookie, convert it to int
	count, err := strconv.Atoi(cookie.Value)
	if err != nil {
		log.Fatalln(err) //we are doing this fatal log, so we can just exit the program and not have to increament the counter!
	}
	count++                            //increment the cookie value
	cookie.Value = strconv.Itoa(count) //convert the counter back to a string

	http.SetCookie(res, cookie) //assign the counter to the cookie value

	io.WriteString(res, cookie.Value) //write back the cookie value to the client
}
