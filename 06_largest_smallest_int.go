//Author: Daire Ní Chatháin
//Problem sheet 1 : Largest and Smallest in Int
/* This Program that returns the largest  and smallest elements in a list.
Adapted from : https://golang.org/pkg/sort/
*/
package main

import "fmt"
import "sort"

func main() {
	s := []int{5, 2, 6, 3, 1, 4} // unsorted
	fmt.Println("Unsorted : ", s)
	sort.Ints(s)
	fmt.Println("Sorted: ", s)
}
