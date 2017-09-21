//Author: Daire Ní Chatháin
//Problem sheet 1 : Factorial Digit Sum
/* A  program that calculates the sum of the digits
 of the factorial of a number 
 */
package main

import "fmt"

func main() {

	var num int  = 10
	var factorial int = 1
	
	for i := 0; i < num; i++{
		
		factorial += factorial * i 
	}

	fmt.Println(factorial)

	var digit, sum int =  0,0;

	for factorial > 0 {
		digit = factorial  % 10
		sum += digit
		factorial /=10
	}
	fmt.Println(sum)
}
