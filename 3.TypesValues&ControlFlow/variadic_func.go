package main

import "fmt"

func read_args(arg ...int) {
	for _, a := range arg {
		fmt.Printf("arg: %d", a)
	}
}

func main() {
	read_args(1, 12, 123, 1234)
}
