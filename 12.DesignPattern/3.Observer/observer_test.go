package observer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestObserver
func TestObserver(t *testing.T) {
	o := &Observer{"test"}
	s := &Subject{}
	s.Add(o)
	s.Notify("Hello")
	assert.Equal(t, "Hello", o.data)
}
