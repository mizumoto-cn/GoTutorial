// ch := make(chan type, n)
// n is 0 -> non-buffer (block)
// n > 0  buffer (non-block untill n elements in channel)
package main

import "fmt"

func fibonacci(n int, ch chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	// remember always just close on the producer side
	// or it's easy to get into panic
	close(ch)
}

func main() {
	c := make(chan int, 2)
	c <- 1
	c <- 2
	fmt.Println(<-c)
	fmt.Println(<-c)
	// operate channels with range
	ch := make(chan int, 10)
	go fibonacci(cap(ch), ch) // cap() Channel: the channel buffer capacity, in units of elements;
	for i := range c {
		fmt.Println(i)
	}
}
