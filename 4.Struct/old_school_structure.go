package main

import (
	f "fmt"
)

type person struct {
	age  int
	name string
}

func main() {
	P0 := person{24, "Mizumoto"}
	P1 := person{age: 23, name: "M"}

	// I love this anonymous way
	P2 := struct {
		age  int
		name string
	}{29, "Airi"}

	var Tod person
	Tod.age = 21
	Tod.name = "Tod"

	f.Println(P0)
	f.Println(P1)
	f.Println(P2)
	f.Println(Tod)
}
