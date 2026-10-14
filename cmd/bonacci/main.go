// Bonacci is a command-line interface for generating a Fibonacci sequence.
//
// This is the placeholder CLI of a project template; remove this directory
// along with the go-arg dependency if your project doesn't need one.
package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/alexflint/go-arg"
	"github.com/leonardo/bonacci"
)

type args struct {
	N int `arg:"positional,required" help:"The number of terms"`
}

func (args) Description() string { // coverage-ignore
	return "Generate a Fibonacci sequence up to the given number of terms"
}

func (args) Version() string { // coverage-ignore
	return "0.1.0"
}

func main() { // coverage-ignore
	var a args
	arg.MustParse(&a)

	sequence := bonacci.FibonacciSequence(a.N)
	terms := make([]string, len(sequence))
	for i, term := range sequence {
		terms[i] = strconv.Itoa(term)
	}
	fmt.Println(strings.Join(terms, " "))
}
