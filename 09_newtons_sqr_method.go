//Author: Daire Ní Chatháin
//Problem sheet 1 : Newtons Square Root Method
// newtons algorithm for guessing the square root of a number
package main

import "fmt"
import "math"

func main() {

	x := 256.0
	z := float64(1)

	for z = 1.0; z != z_next(z, x); z = z_next(z, x) {
		fmt.Printf("Current guess: %40.8f\n", z)
	}

	fmt.Printf("The square root of %.8f is %.8f.\n", x, z)

	fmt.Printf("The math.Sqrt gives the value of %.2f", math.Sqrt(x))
}

func z_next(z float64, x float64) float64 {
	return z - ((z*z - x) / (2 * z))
}
