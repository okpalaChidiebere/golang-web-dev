package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

/*
How we can use context to timeout a process
*/

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	ctx = context.WithValue(ctx, "userID", 777)
	ctx = context.WithValue(ctx, "fname", "Bond")

	results, err := dbAccess(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusRequestTimeout)
		return
	}

	fmt.Fprintln(w, results)
}

func dbAccess(ctx context.Context) (int, error) {

	ctx, cancel := context.WithTimeout(ctx, 1*time.Second) //we basically time the context out after one second
	defer cancel()                                         //we defer the cancel() to avoid additional processes related to this path running to avoid contex leak. The cancel() method is one of the return values by context.WithTimeout()

	ch := make(chan int)

	go func() {
		// ridiculous long running task
		uid := ctx.Value("userID").(int)
		time.Sleep(10 * time.Second)

		// check to make sure we're not running in vain
		// if ctx.Done() has
		if ctx.Err() != nil { //we can check if the context has timed out inside our go routine as well by checking if the context has an error
			return //if the context has an error we stop processiing and return. We leave the go function
		}

		ch <- uid
	}()

	select {
	case <-ctx.Done(): //if the process got canceled as a result of the defer cancel() code at line 40, this will be true
		return 0, ctx.Err()
	case i := <-ch: //at the point, process returned before the context got timed out, so we can get the value returned from the channel. But according to this code, we will not get here because we have waited to long at line 49
		return i, nil
	}
}

func bar(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	log.Println(ctx)
	fmt.Fprintln(w, ctx)
}

// per request variables
// good candidate for putting into context
