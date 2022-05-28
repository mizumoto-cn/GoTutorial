package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

// "A method is a function with an implicit first argument, called a receiver."
// area() is a method somewhat interface like
// func (r ReceiverType) funcName(parameters) (results)
func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func main() {
	r1 := Rectangle{12, 2}
	r2 := Rectangle{9, 4}
	c1 := Circle{10}
	c2 := Circle{25}

	fmt.Println("Area of r1 is: ", r1.area())
	fmt.Println("Area of r2 is: ", r2.area())
	fmt.Println("Area of c1 is: ", c1.area())
	fmt.Println("Area of c2 is: ", c2.area())
}

// If the name of methods are the same but they don't share the same receivers, they are not the same.
// Methods are able to access fields within receivers.
// Use . to call a method in the struct, the same way fields are called.

// go is intellegent enough to interpret so just need to define the receiver as pointer in methods!
// feel free to choose using using &/*'s or not
