// All identifiers defined within a package are visible throughout that package.
// When importing a package you can access only its exported identifiers.
// An identifier is exported if it begins with a capital letter.
package timer

import (
	"fmt"
	"time"
)

type Timestamp struct {
	stamp      time.Time
	is_running bool
}

// Start turns the clock on.
func (s *Timestamp) Set() {
	if !s.is_running {
		s.stamp = time.Now()
		s.is_running = true
		fmt.Println("Timestamp set at ", s.stamp)
	}
}
