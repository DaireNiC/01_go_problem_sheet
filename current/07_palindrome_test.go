// Author: Daire Ní Chatháin
// Problem sheet 1 : Palindrome test
/* This Program allows the user to enter and string and
tests whether that string is a palindrome or not.
Adapted from: https://gist.github.com/renandanton/e99eb16f46e3decceca59b3dbe9f0bb4
*/
package main

import "fmt"

func main() {

	var word string
	var isPalindrome = false

	fmt.Println("=== Palindrome Tester ===")
	fmt.Println("Enter a word: ")

	//User input word
	fmt.Scanln(&word)
	length := len(word)
	//check if each half of word are identical
	for i := 0; i < (length / 2); i++ {
		if string(word[i]) != string(word[length-1-i]) {
			isPalindrome = false
		} else {
			isPalindrome = true
		}
	}

	// Print result
	if isPalindrome {
		fmt.Println("Your word is a palindrome")
	} else {
		fmt.Println("Your word is  not a palindrome")
	}

}
