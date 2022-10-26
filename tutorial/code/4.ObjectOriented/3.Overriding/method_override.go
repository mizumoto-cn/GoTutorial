package main

import "fmt"

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

func (p Hooman) hello() {
	fmt.Printf("Hi I'm %s, I'm %d yo.\n", p.name, p.age)
}

func (p Salaryman) hello() {
	fmt.Printf("Hi I'm %s, I'm %d yo. I work at %s.\n", p.name, p.age, p.company)
}

func main() {
	joe := Salaryman{Hooman{"Joe", 82}, "White House"}
	jimmy := Student{Hooman{"Jimmy", 22}, "Chengdu MDY"}
	joe.hello()
	jimmy.hello()
}
