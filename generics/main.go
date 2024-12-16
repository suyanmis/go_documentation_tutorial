package main

import "fmt"

func main() {
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	floats := map[string]float64{
		"first":  35.94,
		"second": 12.55,
	}

	fmt.Printf("SumFloats(floats): %v\n", SumFloats(floats))
	fmt.Printf("SumInts(ints): %v\n", SumInts(ints))

	fmt.Printf("SumDecimals(floats): %v\n", SumDecimals(floats))
	fmt.Printf("SumDecimals(ints): %v\n", SumDecimals(ints))
}

func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

type Decimal interface {
	int | int32 | int64 | float32 | float64 | int8 | int16
}

func SumDecimals[K comparable, V Decimal](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
