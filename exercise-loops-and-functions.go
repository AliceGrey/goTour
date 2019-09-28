package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {

	z := 1.0

	for i := 0; i < 10; i++ {
		z -=  (z*z - x) / (2*z)
		fmt.Println(z*z)
		if z*z - x <= 0.001 {break}
	}
	return z
}


func main() {
	fmt.Println(Sqrt(2))
}
