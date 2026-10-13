package bonacci

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacciSequence(t *testing.T) {
	assert.Equal(t, []int{}, FibonacciSequence(0))
	assert.Equal(t, []int{1}, FibonacciSequence(1))
	assert.Equal(t, []int{1, 1}, FibonacciSequence(2))
	assert.Equal(t, []int{1, 1, 2, 3, 5}, FibonacciSequence(5))
}
