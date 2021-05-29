package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

/*
This whole code is an example of adding a data to a context and getting the data as well
*/

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	ctx = context.WithValue(ctx, "userID", 777) //this is how you pass a data to context. Remember, dont pass all values to conext. Only important data directly associated with the process like userIDs or sessionIDs
	ctx = context.WithValue(ctx, "fname", "Bond")

	results := dbAccess(ctx)

	fmt.Fprintln(w, results)
}

func dbAccess(ctx context.Context) int {
	uid := ctx.Value("userID").(int) //we get one of the context values. NOTE: .(int) is an assetion technique in go. it thesame like "as" keyword in typescript
	return uid
}

func bar(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	log.Println(ctx)
	fmt.Fprintln(w, ctx)
}

// per request variables
// good candidate for putting into context
