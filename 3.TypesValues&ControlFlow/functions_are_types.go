// type typename func(inputs inputTypes)(results resultTypes)
// useful when writing interfaces
package main

import (
	"fmt"
)

type test func(int) bool

func is_odd(i int) bool {
	return i%2 != 0
}

func is_even(i int) bool {
	return i%2 == 0
}

func filter(slice []int, f test) []int {
	var res []int
	for _, value := range slice {
		if f(value) {
			res = append(res, value)
		}
	}
	return res
}

// A slice is declared just like an array, but it doesn’t contain the size of the slice.
// So it can grow or shrink according to the requirement.
var slice = []int{1, 2, 3, 5, 7, 9}

// more like a pointer with known space and changable length
// slice := make([]T, 5)
// slice := make([]T, 3, 5)
// slice := []int{1, 2, 3, 4} length=space=4
// .append() adds element to the end
// if the known space is full golang will creat a new array and
// copy the former data to the new array
// if the known space < 1000 new array will have 2 times the space
// else 125%

func main() {
	odd := filter(slice, is_odd)
	even := filter(slice, is_even)

	fmt.Println("Slice is ", slice)
	fmt.Println("Odd elements are: ", odd)
	fmt.Println("Even elements are: ", even)
}
