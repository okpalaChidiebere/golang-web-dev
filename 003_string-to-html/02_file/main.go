package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := "Todd McLeod"
	//We merge data strings together using the fmt.Sprint() method
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

	nf, err := os.Create("index.html") //we create a file using the Create() method from the os package
	//Error check to make sure the file is successfully created or not
	if err != nil {
		log.Fatal("error creating file", err)
	}
	defer nf.Close() //close the file

	//Remember io.Copy() takes in the Writer(a standard output; in our case we are writing to a file "nf") and a Reader(were you want to read from)
	io.Copy(nf, strings.NewReader(str))
}
