package strategy

import (
	"testing"
)

// TestLunch tests the Lunch
func TestLunch(t *testing.T) {
	a := NewLunch(&foodCtx{"a", 100}, &Pizza{})
	b := NewLunch(&foodCtx{"b", 200}, &Steak{})
	a.Eat()
	b.Eat()
}
