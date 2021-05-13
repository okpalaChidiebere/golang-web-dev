package main

import (
	"encoding/json"
	"log"
	"net/http"
)

/*
The code here is an example of taking a go datastructure and turning it inot JSON using marshal and encode.
*/

type person struct {
	//FYI these fields needs to be capitalised if you want them to export them to JSON
	Fname string
	Lname string
	Items []string
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/mshl", mshl)
	http.HandleFunc("/encd", encd)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	s := `<!DOCTYPE html>
		<html lang="en">
		<head>
		<meta charset="UTF-8">
		<title>FOO</title>
		</head>
		<body>
		You are at foo
		</body>
		</html>`
	w.Write([]byte(s)) //basic conversion of strings to a slice of bytes
}

func mshl(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p1 := person{
		Fname: "James",
		Lname: "Bond",
		Items: []string{"Suit", "Gun", "Wry sense of humor"},
	}
	j, err := json.Marshal(p1) //the Marshal() method returns a slice of bytes
	if err != nil {
		log.Println(err)
	}
	w.Write(j) //we can write back to the request the bytes!. If we have a writer that took a string we could convert the slice of bytes to a string easily
}

func encd(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p1 := person{
		Fname: "James",
		Lname: "Bond",
		Items: []string{"Suit", "Gun", "Wry sense of humor"},
	}
	err := json.NewEncoder(w).Encode(p1) //we encode to our writer and pass in our go data structure
	if err != nil {
		log.Println(err)
	}
}
