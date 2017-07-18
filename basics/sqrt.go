package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	const accuracy = 0.001
	var lower, upper, guess float64
	if x < 1 {
		lower = x
		upper = 1
	} else {
		lower = 1
		upper = x
	}

	for ((upper - lower) > accuracy) {
		guess = (lower + upper) / 2
		if (guess * guess > x) {
			upper = guess
		} else {
			lower = guess
		}
	}
	return (lower + upper) / 2
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}

