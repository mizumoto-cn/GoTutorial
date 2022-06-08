package main

import (
	"reflect"
)

// this wont work
func swap_false(a, b *interface{}) error {
	*a, *b = *b, *a
	return nil
}

//true

func main() {
	swap := func(in []reflect.Value) []reflect.Value {
		ra := in[0].Elem()
		rb := in[1].Elem()
		tmp := ra.Interface() //var i interface{} = (v's underlying value)

		ra.Set(rb)
		rb.Set(reflect.ValueOf(tmp))

		return nil
	}

	makeSwap := func(fptr interface{}) {
		fn := reflect.ValueOf(fptr).Elem()
		v := reflect.MakeFunc(fn.Type(), swap)
		fn.Set(v)
	}

	a, b := 1, 2
	var intSwap func(*int, *int)

	makeSwap(&intSwap)
	intSwap(&a, &b)
	println(a)
	println(b)
}
