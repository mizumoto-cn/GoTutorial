package main

import (
	. "fmt"
	f "fmt"

	// _ in import: only calls the init() func of the package at first
	// as we might be not sure whether to use the funcs belong to this package
	_ "github.com/ziutek/mymysql/godrv"
)

func init() {
	Println("line not using fmt!")
}

func main() {
	f.Println("line using f as an alias of fmt!")
}
