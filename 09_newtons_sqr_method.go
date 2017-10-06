//Author: Daire Ní Chatháin
//Problem sheet 1 : Newtons Square Root Method
// newtons algorithm for guessing the square root of a number
// Completed in labs

package main

import "fmt"
import "math"

func main() {

	x := 256.0
	z := float64(1) // First guess
	
	// Run code while z doesn't equal the current guess
	for z = 1.0; z != z_next(z, x); z = z_next(z, x) {
		fmt.Printf("Current guess: %40.8f\n", z)
	}
	
	// Print the approx square root value according to Netwon's Method
	fmt.Printf("The square root of %.8f is %.8f.\n", x, z)
	
	// Print the square root calculated by math.Sqrt()
	fmt.Printf("The math.Sqrt gives the value of %.2f", math.Sqrt(x))
}

func z_next(z float64, x float64) float64 {
	//Formula for Netwon's Square Root
	return z - ((z*z - x) / (2 * z))
}
