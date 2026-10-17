package bonacci_test

import (
	"testing"

	"github.com/leonardo/bonacci"
	"github.com/stretchr/testify/assert"
)

func TestFibonacciSequence(t *testing.T) {
	t.Parallel()

	assert.Equal(t, []int{}, bonacci.FibonacciSequence(0))
	assert.Equal(t, []int{1}, bonacci.FibonacciSequence(1))
	assert.Equal(t, []int{1, 1}, bonacci.FibonacciSequence(2))
	assert.Equal(t, []int{1, 1, 2, 3, 5}, bonacci.FibonacciSequence(5))
}
