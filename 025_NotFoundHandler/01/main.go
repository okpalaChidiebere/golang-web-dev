package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	/*The browser like google chrome will often ask you for your favicon.ico. SOme browser don't ask for it
	faveicon.com is the little image that shows on the tab of the browser for the site page you are on
	If you dont have the faveicon you want to be able to respond back to the browser request that you dont have one with like a 404 code

	So, in our case, we passed the NotFoundHandler() method which does that for us.

	http.NotFoundHandler() return a Handler thats why we can pass it to the Handle() method
	https://golang.org/pkg/net/http/#NotFoundHandler
	*/
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL.Path)
	fmt.Fprintln(w, "go look at your terminal")
}
