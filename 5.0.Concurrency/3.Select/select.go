// select is blocking by default
// you can listen to loads of channels using select
// if there are multiple channels avaliable
// it will randomly choose one to execute
package main

import "fmt"

func fibonacci(ch, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case ch <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

// select also has "default"
// becomes non-blocking
// do something default and not wait for channel more

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	fibonacci(c, quit)
}
