package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080") //on the web we use 80 but due to we are testing we use 8080
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err.Error())
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	// read request
	request(conn)

	// write response
	respond(conn)
}

func request(conn net.Conn) {
	i := 0 //we use this i variable a flag in our program
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		/*
			if the flag is 0 it means we are in the first iteration of the loop which
			means its the first time we are reading this request. This also means we
			just read the requestLine which is the first line for every HTTP request
			according to RFC7230

			Rememeber the scanner.Text() gets the scans the request in text form line by line
		*/
		if i == 0 {
			// request line
			m := strings.Fields(ln)[0] // strings.Fields() method breaks a string into words. So now we can get the first word in the requestLine which is the METHOD for the request. As we can parse the method(eg GET, POST, PATCH) and URI(eg / or /invoice/add) out the requestLine we can start to make conditional logic to run some certain block of code based on these values
			fmt.Println("***METHOD", m)
			fmt.Println("***URI", strings.Fields(ln)[1]) //gets the URI or path
		}
		/*
			if line is equal to nothing we want to break out of the loop

			Remember that accoridng to the spec in rfc7230, we have a blank line between the
			requestLine, Header and body then a blank line at the end of the request as well
			We just read the header and we can break out the loop. For your own server, you possibly will want to read the body as well
		*/
		if ln == "" {
			// headers are done
			break
		}
		i++
	}
}

func respond(conn net.Conn) {

	//We created a string of text for the response
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>Hello World</strong></body></html>`

	/*
		We are printing back our response to the client and we have to format our response to adhere to the HTTP Protocol rfc7230
	*/
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")                //first our startLine
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body)) //Content-Length is a header you should send back to browser
	fmt.Fprint(conn, "Content-Type: text/html\r\n")        //content-type is a html. in your project is the response is a json from database you can put that here
	fmt.Fprint(conn, "\r\n")                               //\r\n means a blank line
	fmt.Fprint(conn, body)
}
