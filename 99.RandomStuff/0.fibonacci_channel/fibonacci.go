package main

import (
	"fmt"
)

func fibonacci(n int) int {
	switch n {
	case 1:
		return 1
	case 2:
		return 1
	default:
		return fibonacci(n-1) + fibonacci(n-2)
	}
}

func fibonacci_list(n int) {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibocon(c, quit)
}

func fibocon(c, quit chan int) {
	x, y := 1, 1
	for {
		// select always randomly selects one, if one is blocked then choose another
		// if none of them is avaliable block and wait
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	println(fibonacci(10))
	fibonacci_list(10)
}
