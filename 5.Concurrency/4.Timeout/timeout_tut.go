package main

import "time"

func main() {
	c := make(chan int)
	o := make(chan bool)

	go func() {
	L:
		for {
			println("beep")
			select {
			case v := <-c:
				println(v)
			case <-time.After(5 * time.Second):
				println("timeout")
				o <- true
				break L
			}
		}
	}()
	println(<-o)
}
