`runtime.Goexit()`

Exits the current goroutine, but defered functions will be executed as usual.

`runtime.Gosched()`

Lets the scheduler execute other goroutines and comes back at some point.

`runtime.NumCPU() int`

Returns the number of CPU cores

`runtime.NumGoroutine() int`

Returns the number of goroutines

`runtime.GOMAXPROCS(n int) int`

Sets how many CPU cores you want to use