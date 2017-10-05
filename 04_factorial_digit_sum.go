//Author: Daire Ní Chatháin
//Problem sheet 1 : Factorial Digit Sum
/* A  program that calculates factorial of a number & the sum of the digits
 of that number
 */
package main

import (
	"fmt"
	"math/big"
)

func main() {

	//var num int  = 10
	 num := big.NewInt(10)
	//var factorial int = 1
	factorial:= big.NewInt(1)
	i := big.NewInt(0)
	
	for  i < num; i++{
		//calculate factorial
		factorial += factorial * i 
	}
	//print factorial
	fmt.Println(factorial)

	//var digit, sum int =  0,0;
	digit:= big.NewInt(0)
	sum:= big.NewInt(0)
	
	
	//while there are  are more than 0 digits in the number
	for factorial > 0 {
		//split num by using modulo to separate digits
		digit = factorial  % 10
		//add digit to sum
		sum += digit
		//update new factorial value (one digit less each time)
		factorial /=10
	}
	fmt.Println(sum)
}
