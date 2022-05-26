// if...else...
// goto  case sensitive
// for break
// for expa ; expb ; expc {} just like c without '()'
// parallel assignment    go doesnt have , operator so it is like a, b = c, d
// for in range ( array slice map string)
// for k, v := range map { }   when you dont wwant a return value use _ instead
// "while(true)" be like for { ... }
// same switch ... case ... like C but not jump out by "break" instead it will jump out by default
// so you may need "fallthrough" to match more cases below
// func funcName (input input type...) (return type...) { function body ... return ...}
//

package main

// #include <math.h>
// #define max(a,b) (a)>(b)?a:b
// int Cmax (int a, int b){
//	return max(a, b);
//}
import "C"
import (
	"fmt"
)

func go_max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	x, y, z := 3, 4, 5
	gomax_xy := go_max(x, y)
	gomax_yz := go_max(y, z)
	cmax_xy := C.Cmax(C.int(x), C.int(y)) // y'all have to use
	cmax_yz := C.Cmax(C.int(y), C.int(z))
	fmt.Printf("max %d %d in go is %d in C is %d", x, y, gomax_xy, cmax_xy)
	fmt.Printf("max %d %d in go is %d in C is %d", x, z, gomax_yz, cmax_yz)
}
