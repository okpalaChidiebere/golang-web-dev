package main

import "fmt"

func main() {
	name := "Todd McLeod"

	//We have a string that we merge to a variable
	tpl := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Hello World!</title>
	</head>
	<body>
	<h1>` + name + `</h1>
	</body>
	</html>
	`
	fmt.Println(tpl) //the we print the concatenated string out in terminal

	//We can go on dump the printed result to standardOut(stdout), in our case its a file; through a pipeline with go CLI
	//eg go run main.go > index.html
	//instead of printing to commandline we print the result to a html file!
}
