package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	//We loop forever to wait or listen for an incomming connection
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		//We launch a go routine for our function that accepts a connection as argument
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	/*
		We use bufio to read from the incomming connection

		NewScanner method returns a pointer to a scanner (*scanner), so you can now call
		the methods that accepts the pointer scanner as receivers

		https://golang.org/pkg/bufio/#Scanner
	*/
	scanner := bufio.NewScanner(conn) //to see how the scanner works and scan through string text go to 000_temp/26_bufio-NewScanner/...
	for scanner.Scan() {              //Rememeber the Scan method returns a bool. If true means we have not reached the end of input
		ln := scanner.Text() //We scan the text line by line by default
		fmt.Println(ln)      //Print the line
	}
	defer conn.Close()

	// we never get here
	// we have an open stream connection
	// how does the above reader know when it's done?
	/*
		We never get to this code because, knowing that the Scan() method returns
		false only when it encounters an error or it has reached end of input but we have a wide open stream
		So our scanner thinks that more text is coming so it just keep looping and never gets to defer conn.Close()

		Your Scanner by default dont know when to stop scanning the open connection. We have to tell it ourself! We
		will fix this in the next program


	*/
	fmt.Println("Code got here.")
}
