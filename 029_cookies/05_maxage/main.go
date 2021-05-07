package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index) //loads the home page with a link to the 'set' route
	http.HandleFunc("/set", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/expire", expire)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, `<h1><a href="/set">set a cookie</a></h1>`)
}

func set(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "session",
		Value: "some value",
		Path:  "/",
	})
	fmt.Fprintln(w, `<h1><a href="/read">read</a></h1>`)
}

func read(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session")
	if err != nil { //at this point, it means there is no cookie
		http.Redirect(w, req, "/set", http.StatusSeeOther) //we just redirect the user back to the set route. Yey! here we go again with 303
		return                                             //we dont forget to include the return statement
	}

	fmt.Fprintf(w, `<h1>Your Cookie:<br>%v</h1><h1><a href="/expire">expire</a></h1>`, c)
}

func expire(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session")
	if err != nil {
		http.Redirect(w, req, "/set", http.StatusSeeOther)
		return //we must return, so that we dont go ahead to delete a cookie when there is no cookie!
	}
	c.MaxAge = -1                                   // delete cookie. NOTE: You can set the age to 0 or a negative value. They are all thesame. It will dlete the cookie
	http.SetCookie(w, c)                            //we need to seek the cookie as we hav change one of the fields
	http.Redirect(w, req, "/", http.StatusSeeOther) //we redirect back to the default route
}
