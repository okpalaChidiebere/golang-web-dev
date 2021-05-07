package main

import (
	"html/template"
	//"github.com/satori/go.uuid"
	"net/http"
)

type user struct {
	UserName string
	First    string
	Last     string
}

var tpl *template.Template
var dbUsers = map[string]user{}      // user ID, user
var dbSessions = map[string]string{} // session ID, user ID
//an aletrnative way we could have initialize an empty map is to use the make method. eg make(map[string]string) instead of map[string]string{}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {

	// get cookie
	c, err := req.Cookie("session") //we ask for a cookie
	if err != nil {                 //if the cookie is not there(set in the client machine, and sent with this request), we create one
		sID, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
	}

	// if the user exists already, get user
	var u user
	/*
		c.Value s the value of the cookie, but for the map, the cookie value is
		the key that will give us the username.
		The username according to our app must be unique. But ideally we will user a userID instead of name

		The username is now a key for the users map table "dbUsers"
	*/
	if un, ok := dbSessions[c.Value]; ok {
		u = dbUsers[un] //  we now get the username of the user by id from the user table
	}

	// process form submission
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		u = user{un, f, l}       //we create a new user
		dbSessions[c.Value] = un //we associate the sessionId with the username
		dbUsers[un] = u
	}

	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

/*
With this function, if there is a session, we display the username of the user logged in

Remember the username is associated to a session in this app
*/
func bar(w http.ResponseWriter, req *http.Request) {

	// get cookie
	c, err := req.Cookie("session")
	if err != nil {
		http.Redirect(w, req, "/", http.StatusSeeOther) //if there is no session, we just redirect back to the default route and not print anything
		return
	}
	un, ok := dbSessions[c.Value]
	/*
		If the user session has a username attached to it(something we stored),
		ok value will be true and we will not run the redirect code
	*/
	if !ok {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	u := dbUsers[un]
	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}

// map examples with the comma, ok idiom
// https://play.golang.org/p/OKGL6phY_x
// https://play.golang.org/p/yORyGUZviV
