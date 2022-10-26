package main

import (
	"fmt"
)

type Hooman struct {
	name string
	age  int
}

type Student struct {
	Hooman // anonymous field
	school string
}

type Salaryman struct {
	Hooman
	company string
}

func (p *Hooman) poopoo() {
	fmt.Printf("%s, %d years old, shitted his pants!\n", p.name, p.age)
}

func main() {
	yjspi := Student{Hooman{"Tadokoro", 24}, "ChuntianhuahuaYouzhiyuan"}
	joe := Salaryman{Hooman{"Joe", 82}, "White house"}

	yjspi.poopoo()
	joe.poopoo()
}
