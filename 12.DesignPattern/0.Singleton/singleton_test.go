package singleton

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestGetInstance tests the GetInstance function.
func TestGetInstance(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(100)
	start_when_close := make(chan struct{})
	array := make([]*Singleton, 100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			defer wg.Done()
			// block until the start singal
			<-start_when_close
			array[i] = GetInstance()
		}(i)
	}
	close(start_when_close)
	wg.Wait()

	for i := 0; i < 100; i++ {
		assert.Equal(t, array[0], array[i])
	}
}
