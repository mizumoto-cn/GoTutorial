package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("Hello from Go Version %s", runtime.Version())
}
