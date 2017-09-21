//Author: Daire Ní Chatháin
//Problem sheet 1 : Palindrome test
/* This Program that tests whether a string is a palindrome.
Adapted from: 
*/
package main

import "fmt"

func main() {

	var word string = ""
	var isPalindrome = False
	fmt.Println("=== Palindrome Tester ===")
	fmt.Println("Enter a word: ")

	//input of word

    fmt.Scanln(&word)
	length := len(word)

	for i:=0 ; i < (length/2) ; i++ {
		if string(word[i]) != string(word[length - i -1]){

		}else{
			isPalindrome = true
		}
	
	}

}