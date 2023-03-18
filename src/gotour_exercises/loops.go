package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := x / 2.0
	z_prev := x + 1
	for i := 0; math.Abs(z - z_prev) > 0.000000000000001; i++ {
		z_prev = z
		z -= (z*z - x) / (2*z)
		fmt.Println(z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}