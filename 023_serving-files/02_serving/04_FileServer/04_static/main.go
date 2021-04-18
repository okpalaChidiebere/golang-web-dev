package main

import (
	"log"
	"net/http"
)

func main() {
	/*
		http.ListenAndServe returns an error value

		if ListenAndServe fails and throws an error, that error gets returned, main will exit and our program will be done.

		However, ideally, we wil want to catch that error and log it.
		So log.Fatal accepts any type (represented like ...interface{}) as argument and prints out
	*/
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("."))))
}

/*
The code here is thesame to

    http.Handle("/", http.FileServer(http.Dir(".")))
	http.ListenAndServe(":8080", nil)


NOTE:
    Becasue we have a file called index.html, it will no longer show our folders or files when we go to root (/)
	It says that in the documentation https://golang.org/pkg/net/http/#FileServer
	So index.html is a MUST have to serve your static website successfully
*/
