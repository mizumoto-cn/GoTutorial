package main

import "fmt"

func sum_and_product(a, b int) (int, int) {
	return a + b, a * b
}

func main() {
	a, b := 2, 3
	aplusb, atimesb := sum_and_product(a, b)
	fmt.Printf("%d + %d = %d, %d * %d = %d", a, b, aplusb, a, b, atimesb)
}

// also be like
// func sum_and_product(a, b int) (sum int, product int) {
// 	sum = a + b
// 	product = a * b
// 	return
// }
