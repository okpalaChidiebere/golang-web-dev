package main

import (
	"log"
	"net/http"
)

func main() {
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
