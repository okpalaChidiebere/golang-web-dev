package main

import (
	"fmt"
	"net/http"
)

type hotdog int

//We know hotdog is a Handler type as well
func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//Inside this function, I can write back as respnse any code you want
	//But for this program we just print this to the web when the user goes to http://localhost:8080/
	fmt.Fprintln(w, "Any code you want in this func")
}

func main() {
	var d hotdog                    //create a variable of type hotdog
	http.ListenAndServe(":8080", d) //pass thatvariable to ListenAndServe. This is not a HTTPS connection. If you want HTTPS look at the README file
}
