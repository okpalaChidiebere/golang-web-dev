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

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text() //scan whenever the user typed in. Immediately the user enters the "\n" symbol by pressing Enter in terminal in the connection
		fmt.Println(ln)
		fmt.Fprintf(conn, "I heard you say: %s\n", ln) //We write back to the user whatever we read in :)
	}
	defer conn.Close()

	// we never get here
	// we have an open stream connection
	// how does the above reader know when it's done?
	//still we have not fixed the closing of connection
	fmt.Println("Code got here.")
}
