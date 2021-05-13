package main

import (
	"fmt"
	"time"
)

/*
This whole code is broken
*/

func main() {
	for n := range gen() {
		fmt.Println(n)

		/*
			This is the problem with this code.
			After the 5 value is return from the channel, we want to break out the loop. But it will not
			because the gen() method is broken as well. The gen() method runs forever and keep returning values to the channel to our range loop

			Although there is a break, it will not work, because our range is attached to that channel. It is not a regular for loop that runs on a counter

			So we need a better way to break out the loop after 5. We want to stop the go routine after the 5th value is returned so this loop can work properly as well
		*/
		if n == 5 {
			break
		}
	}
	time.Sleep(1 * time.Minute)
}

// gen is a broken generator that will leak a goroutine.
func gen() <-chan int {
	ch := make(chan int)
	go func() {
		var n int
		for {
			ch <- n
			n++
		}
	}()
	return ch
}
