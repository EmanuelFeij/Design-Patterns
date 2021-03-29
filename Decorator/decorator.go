package Decorator

import (
	"fmt"
	"math"
)

func Square(x float64) float64 {
	return math.Sqrt(x)
}

func IsGreaterThanZero(x float64) bool {
	return x >= 0
}

func Sum3(x float64) float64 {
	return x + 3
}

// Here can pass any function that matches the argument
func DoThingsWithPositiveFloat(fn func(float64) float64, x float64) float64 {
	if !IsGreaterThanZero(x) {
		return 0
	}

	return fn(x)
}

func main() {
	squared := DoThingsWithPositiveFloat(Square, 4.0)
	plus3 := DoThingsWithPositiveFloat(Sum3, squared)
	fmt.Println(plus3)
}
