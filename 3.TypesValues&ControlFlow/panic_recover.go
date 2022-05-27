// panic recover system is powerful but you'd better not us them taht much
package main

import "fmt"

// Panic is a built-in function to break the current control flow
// when a func 'F' calls panic, F will stop executing but the
// defer func will continue to execute
// then 'F' returns to the break point which caused the panic
// Program waits until all functions return with panic to the
// first level of goroutine
// Panic happens when panic is called or error happenes
//
// Recover is another built-in func to recover goroutine from panic
// calling recover in defer funcs cathces panic values and go back to normal flow
// when it's not in panic status, recover returns nil and do nothing

var money int

func check_wallet() {
	if money < 0 {
		panic("bankrupt")
	}
}

func pay(m int) {
	money = money - m
}

func throws_panic(f func()) (b bool) {
	defer func() {
		if x := recover(); x != nil {
			b = true
		}
	}()

	// execute f() to see whether there is panic
	// and possibly recover it
	f()
	return
}

func main() {
	money = 500
	pay(1000)
	fmt.Println(money)
	fmt.Println(throws_panic(check_wallet))
}
