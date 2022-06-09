// to communicate between threads
// is like a two-way pipe in UNIX
// the only type used in channels is
// type channel keyword chan, use make to create new channel
// use <- to send or receive data
// ci := make(chan int)
// ch <- v
// v := <- ch
package main

import "fmt"

func add_up(ai []int, c chan int) {
	total := 0
	for _, v := range ai {
		total += v
	}
	c <- total
}

func main() {
	a := []int{2, 43, 42, 97, 6, 8, 0, 13}
	c := make(chan int)
	go add_up(a[:len(a)/2], c)
	go add_up(a[len(a)/2:], c)
	x, y := <-c, <-c
	// channels are blocked by default, no need to lock
	// goroutine wont continue if receiving data from a empty channel
	// nor will they if they send data to channel, until the data get received
	fmt.Println(x, " ", y, " ", x+y)
}
