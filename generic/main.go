package main

import "fmt"

// SumInts adds together the value of m.
func SumInts(m map[string]int64) (s int64) {
	for _, v := range m {
		s += v
	}
	return
}

// SumFloats add together the values of m.
func SumFloats(m map[string]float64) (s float64) {
	for _, v := range m {
		s += v
	}
	return
}

// SumIntsOfFloats sums the values of map m. Ti supports both int64 and float64 types
// as types for map values.
func SumIntsOrFloats[K comparable, V Number](m map[K]V) (s V) {
	for _, v := range m {
		s += v
	}
	return
}

type any interface{}
type Number interface {
	int64 | float64
}

func main() {
	// Initialize a map for the integer values.
	ints := map[string]int64{
		"first":  1,
		"second": 2,
	}

	// Initialize a map for the float64 values.
	floats := map[string]float64{
		"first":  1.0,
		"second": 2.0,
	}

	fmt.Printf("Non-Generic Sums: %v and %v\n",
		SumInts(ints), SumFloats(floats))

	fmt.Printf("Non-Generic Sums: %v and %v\n",
		SumIntsOrFloats(ints), SumIntsOrFloats(floats))
}
