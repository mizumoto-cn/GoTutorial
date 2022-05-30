// Reference: The Laws od Reflection <https://go.dev/blog/laws-of-reflection>
// Using reflect package go can tell runtime information
package main

// 1. reflect type
// Reflection goes from reflection object to interface value.

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4
	// TypeOf returns the reflection Type of the value in the interface{}.
	//func TypeOf(i interface{}) Type
	fmt.Println("type:", reflect.TypeOf(x))
	//When we call reflect.TypeOf(x), x is first stored in an empty interface,
	//which is then passed as the argument;
	//reflect.TypeOf unpacks that empty interface to recover the type information.
	fmt.Println("value:", reflect.ValueOf(x).String())

	// also remember to write v:=reflect.ValueOf(&x)
	// if you want to use v.SetFloat(1.23)
}
