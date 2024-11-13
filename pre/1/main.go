package main

import "fmt"

// Pair represents a pair of values of different types.
type Pair struct {
	First  interface{}
	Second interface{}
}

// NewPair creates a new Pair from two values of potentially different types.
func NewPair(first interface{}, second interface{}) Pair {
	return Pair{First: first, Second: second}
}

func main() {
	// Creating a pair of int and string
	p1 := NewPair(1, "one")
	fmt.Printf("Pair 1: (%v, %v)\n", p1.First, p1.Second)

	// Creating a pair of float64 and bool
	p2 := NewPair(3.14, true)
	fmt.Printf("Pair 2: (%v, %v)\n", p2.First, p2.Second)
}
