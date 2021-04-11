package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	//You can accept arguments from the commandline
	name := os.Args[1]
	fmt.Println(os.Args[0]) //we dont always care abut the first argument because its just the program we have executed
	fmt.Println(os.Args[1]) //we pass the name variable from the commandline
	//concatenate the strings
	str := fmt.Sprint(`
		<!DOCTYPE html>
		<html lang="en">
		<head>
		<meta charset="UTF-8">
		<title>Hello World!</title>
		</head>
		<body>
		<h1>` +
		name +
		`</h1>
		</body>
		</html>
	`)

	//create a file
	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal("error creating file", err)
	}
	defer nf.Close()

	//dump the html in the file we created!
	io.Copy(nf, strings.NewReader(str))
}

/*
at the terminal:
go run main.go Chidiebere
*/
