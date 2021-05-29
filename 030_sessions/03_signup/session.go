package main

import (
	"net/http"
)

/*
You can distribute the code that is thesame package across multiple files
as long as those files are in thesame folder.

Noteice that we are directly referecing variables declared in main.go here like dbSessions and dbUsers.
We are able to do so because they are under thesame package name
*/

func getUser(req *http.Request) user {
	var u user

	// get cookie
	c, err := req.Cookie("session")
	if err != nil {
		return u
	}

	// if the user exists already, get user
	if un, ok := dbSessions[c.Value]; ok {
		u = dbUsers[un]
	}
	return u
}

func alreadyLoggedIn(req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}
	un := dbSessions[c.Value]
	_, ok := dbUsers[un]
	return ok
}
