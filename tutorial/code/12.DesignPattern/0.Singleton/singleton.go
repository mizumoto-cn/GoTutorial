// When you only need one instance in your application
// the Singleton pattern is a good choice.
// With sunc.Once, you can guarantee that only one instance will be created.
package singleton

import "sync"

var s *Singleton
var once sync.Once

type Singleton struct {
}

func GetInstance() *Singleton {
	// Do calls the function f if and only if Do is being called for the first time for this instance of Once.
	once.Do(func() {
		s = &Singleton{}
	})
	return s
}
