package main

import (
	"fmt"
	"strconv"
)

type Hooman struct {
	name string
	age  int
}

// Human implements fmt.Stringer
func (h Hooman) String() string {
	return "Name: " + h.name + ", Age: " + strconv.Itoa(h.age) + " yo."
}

func main() {
	Bob := Hooman{"Bob", 39}
	fmt.Println("This Hooman: ", Bob)
}
