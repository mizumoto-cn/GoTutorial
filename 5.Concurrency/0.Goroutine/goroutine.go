// goroutines are similar to threads but smaller
// go supports full memory sharing between goroutines
// lightweight efficient convenient
// go keyword to run a new goroutine
// basically a normal function at underlying level,
// managed by a thread manager provided by go runtime
// main() is a goroutine itself
package main

import (
	"fmt"
	"runtime"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		//Gosched yields the processor, allowing other goroutines to run.
		//It does not suspend the current goroutine, so execution resumes automatically.
		runtime.Gosched()

		fmt.Println(s, i)
	}
}

func main() {
	go say("world")
	say("hello")
}
