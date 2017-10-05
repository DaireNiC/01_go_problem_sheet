//Author: Daire Ní Chatháin
//Problem sheet 1 : String reversal

/* A  program that takes input of a string and then reverses it.

Resource: https://stackoverflow.com/questions/1752414/how-to-reverse-a-string-in-go

A rune is a 32 bit value value mapped to it's unicode codepoint. E.g 'a' = 97
Generally, a string is just considered a collection of runes

The assumption with this program is that the input is  an ASCII UTF-8 encoded string
*/

package main

import "fmt"

func main() {
	//Input string for reversing
	fmt.Println(Reverse("The quick brown fox jumped over the lazy dog"))
}

//function takes param of string & returns reversed string
func Reverse(s string) string {

	//get length of string
	n := len(s)

	//Create an array of runes(basically to chars) the lengnth of string
	runes := make([]rune, n)

	//loop through and add each rune to runes array
	//range on strings iterates over Unicode code points.
	for _, rune := range s {
		n--
		runes[n] = rune
	}
	//return reversed string
	return string(runes)
}
