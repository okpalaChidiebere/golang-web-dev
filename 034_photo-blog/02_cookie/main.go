package main

import (
	"html/template"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	c := getCookie(w, req)
	tpl.ExecuteTemplate(w, "index.gohtml", c.Value)
}

// add func to get cookie
//This function returns an existing cookie if it exists or issue the client a new cookie, if you dont have a session cookie and return the new cookie created
func getCookie(w http.ResponseWriter, req *http.Request) *http.Cookie {
	c, err := req.Cookie("session") //we try to get the session cookie we give
	if err != nil {                 //if we have not issued you a cookie or your cookie has expired or was deleted we give you a new cookie
		sID, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
	}
	return c
}
