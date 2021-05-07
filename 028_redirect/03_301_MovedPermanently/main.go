package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	fmt.Print("Your request method at foo: ", req.Method, "\n\n")
}

func bar(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at bar:", req.Method)
	http.Redirect(w, req, "/", http.StatusMovedPermanently)
}

/*

We first to to '/bar' GET method, then we redirect to '/' which runs the foo() method

The intersting thing now is, the next time we try o access the '/bar' path, we will just go straight to '/'

We will not hit the '/bar' location anymore because we indicated that it now permamently moved to '/'
What happens is that your browser remembers that /bar is now moved permanently! This only stops when you clear your browser cache
*/
