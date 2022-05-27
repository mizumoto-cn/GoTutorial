package main

type Hooman struct {
	name   string
	age    int
	weight int
}

type student struct {
	Hooman
	grade int
}

// just let the fucking testing machine shut up
func init() {
	me := student{Hooman{"me", 12, 34}, 6}
	println(me)
}

//not only structs but you can also add built-in types like
// skills (from type skills []string )
// int   (be like xxx.int = 7)
// into structs
