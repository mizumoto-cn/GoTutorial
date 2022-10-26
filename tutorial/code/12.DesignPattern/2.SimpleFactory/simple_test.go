package simpleFactory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestCreate
func TestCreate(t *testing.T) {
	factory := &Factory{}
	pizza := factory.Create("pizza")
	steak := factory.Create("steak")
	assert.Equal(t, "cooked pizza", pizza.create())
	assert.Equal(t, "cooked steak", steak.create())
}
