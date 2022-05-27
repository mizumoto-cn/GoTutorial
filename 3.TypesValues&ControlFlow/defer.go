// defer statements exceutes in reverse order
package main

import (
	"fmt"
)

func print_defers(n int) {
	for i := 0; i < n; i++ {
		defer fmt.Println(i)
	}
}

func fake_file_IO() {
	fmt.Println("Fake file.Open(\"file\")")
	// will automatically execute at the end of the function
	defer fmt.Println("Fake file.Close")
	fmt.Println("Fake File Operate")
}

func main() {
	fake_file_IO()
	print_defers(5)
}
