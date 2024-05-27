package lb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelector(t *testing.T) {
	selector := NewSelector()
	assert.Equal(t, uint64(1), selector.Next())
	assert.Equal(t, uint64(2), selector.Next())
}
