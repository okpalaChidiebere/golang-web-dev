package main

import (
	"context"
	"fmt"
	"time"
)

/*
We fixed the broken code(at ../04) here by using context so we dont have to keep waiting for a leaked go routine!

The code wil just timeout after a while!
*/

func main() {
	ctx, cancel := context.WithCancel(context.Background()) //context background is the process we start off with. context.WithCancel() gives us a new cancel
	defer cancel()                                          // make sure all paths cancel the context to avoid context leak

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			/*
				Now we call cancel() right before we break.
				This will get the gen() process in the background to stop sending data back that we are done
			*/
			cancel() //right before we break
			break
		}
	}
	time.Sleep(1 * time.Minute)
}

func gen(ctx context.Context) <-chan int {
	ch := make(chan int)
	go func() {
		var n int
		for {
			select {
			case <-ctx.Done(): //if the context is done, we stop. calling the cancel() method before we break in the main func wil make this true
				return // avoid leaking of this goroutine when ctx is done.
			case ch <- n:
				n++
			}
		}
	}()
	return ch
}
