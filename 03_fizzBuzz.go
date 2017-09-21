//Author: Daire Ní Chatháin
//Problem sheet 1 : FizzBuzz
/* This program prints the numbers from 1 to 100, except for
the following conditions. For multiples of three prints "Fizz"
instead of the number, and for the multiples of five prints"Buzz".
For numbers which are multiples of both three and five print 
 */
package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++{
		
		var result string = ""

		if i % 3 == 0 {
			result += "fizz"
		}
		if i % 5 == 0 {
			result += "buzz"
		}
		if result != "" {
			fmt.Println(result)
		}else{
			fmt.Println(i)
			continue
		}
		
	}
}
