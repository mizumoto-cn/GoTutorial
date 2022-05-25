package main

import (
	"fmt"
)

// Constants can be used for enumerations.
// Here it shows how the iterator adds to itself
// at first iota gives 0, whenever iota is used again on a new line
// its value is incremented by 1
// when you dont give out a value, it will automaticly copy the former one
const (
	a = iota // a = 0
	b        // b = 1
	c        // c = 2
	d = 5    // d = 5
	e        // e = 5
)

func main() {
	fmt.Printf("a=%d,b=%d,c=%d,d=%d,e=%d", a, b, c, d, e)
}
