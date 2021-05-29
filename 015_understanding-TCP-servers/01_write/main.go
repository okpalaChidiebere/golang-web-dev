package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	/*
		The listen method accept two arguments
		- First, what type of network you want to listen on. In our case the TCP
		- Second, what port you want to listen on
		then returns a listener or an error

		The Listener is an interface that implements the Close() method and the Accept() method
		whuch enbales you to read to a connection or write to a connection
	*/
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	//We loop forever
	for {
		/*
			inside our loop, if someone calls in, we accept!

			The accept method returns a connection or an err
			The connection implements the Reader and Writer interface

			I hope by now, you understand when you see the work like "implements the bla bla bla interface"
		*/
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		/*We write data to our connection */
		io.WriteString(conn, "\nHello from TCP server\n") //WriteString methods accepts two argumments: the Writer(we pass the conn in there because it implements the writer interface!) and a string. This method writes a string to a writer
		fmt.Fprintln(conn, "How is your day?")
		fmt.Fprintf(conn, "%v", "Well, I hope!")
		/*End writing data to our connection */

		conn.Close() //close the connection so that our server will be ready to accept another one
	}
}
