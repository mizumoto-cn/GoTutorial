package main

import (
	"fmt"
	"sync"
)

func main() {
	letter, number := make(chan bool), make(chan bool)
	// To wait for multiple goroutines to finish, we can use a wait group.
	wait := sync.WaitGroup{}
	go func() {
		i := 1
		// Never ever use for { select case}
		for {
			<-number
			fmt.Print(i)
			i++
			fmt.Print(i)
			i++
			letter <- true

		}
	}()
	wait.Add(1)
	go func(wait *sync.WaitGroup) {
		i := 'A'
		for {
			<-letter
			if i >= 'Z' {
				wait.Done()
				return
			}
			fmt.Print(string(i))
			i++
			fmt.Print(string(i))
			i++
			number <- true

		}
	}(&wait)
	number <- true
	wait.Wait()
}
