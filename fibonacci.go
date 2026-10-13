// Package bonacci is an example library for generating a Fibonacci sequence.
//
// This is the placeholder library of a project template; replace
// [FibonacciSequence] with your own logic.
package bonacci

import "slices"

// FibonacciSequence generates a Fibonacci sequence up to the given number of terms.
//
// The sequence starts from 1, 1, and exactly n terms are returned, so an n of
// 0 yields an empty sequence.
func FibonacciSequence(n int) []int {
	sequence := slices.Repeat([]int{1}, n)
	for i := 2; i < n; i++ {
		sequence[i] = sequence[i-2] + sequence[i-1]
	}
	return sequence
}
