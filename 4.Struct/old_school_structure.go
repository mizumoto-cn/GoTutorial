package main

import (
	. "fmt"
)

type person struct {
	age  int
	name string
}

func main() {
	P0 := person{24, "Mizumoto"}
	P1 := person{age: 23, name: "M"}
	P2 := struct {
		age  int
		name string
	}{29, "Airi"}
	var Tod person
	Tod.age = 21
	Tod.name = "Tod"
	Println(P0)
	Println(P1)
	Println(P2)
	Println(Tod)
}
