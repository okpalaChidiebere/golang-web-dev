package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
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
	defer conn.Close()

	// instructions
	io.WriteString(conn, "\r\nIN-MEMORY DATABASE\r\n\r\n"+
		"USE:\r\n"+
		"\tSET key value \r\n"+
		"\tGET key \r\n"+
		"\tDEL key \r\n\r\n"+
		"EXAMPLE:\r\n"+
		"\tSET fav chocolate \r\n"+
		"\tGET fav \r\n\r\n\r\n")

	// read & write
	data := make(map[string]string) // we basically used a map to store our data in memory
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fs := strings.Fields(ln) //string.Fields() method breaks the text comming in a request into individual words and store the words as a slice of strings
		// logic
		if len(fs) < 1 { //the length is not equal to three, we skip the the current connection and its request. We basically will not go further to process the request and we wait for thext connection
			continue
		}
		switch fs[0] {
		case "GET": //eg of a request GET cat   "cat" is the key here
			k := fs[1]
			v := data[k]
			fmt.Fprintf(conn, "%s\r\n", v)
		case "SET": //eg SET cat james    "cat" is the key here and "james" is the value
			if len(fs) != 3 {
				fmt.Fprintln(conn, "EXPECTED VALUE\r\n")
				continue
			}
			k := fs[1]
			v := fs[2]
			data[k] = v //storing the cat vlaue in the map by its key
		case "DEL":
			k := fs[1]
			delete(data, k)
		default:
			fmt.Fprintln(conn, "INVALID COMMAND "+fs[0]+"\r\n")
			continue
		}
	}
}
