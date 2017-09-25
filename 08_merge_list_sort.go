//Author: Daire Ní Chatháin
//Problem sheet 1 : Merge Sort two ordered slices
/* This Program that returns the largest  and smallest elements in a list.
Adapted from : http://austingwalters.com/merge-sort-in-go-golang/
*/
package main

import (
	"fmt"
)

func main() {
	// Merges left and right slice into newly created slice
	left := []int{1, 4, 6}
	right := []int{2, 3, 5}

	size, i, j := len(left)+len(right), 0, 0
	slice := make([]int, size, size)

	for k := 0; k < size; k++ {
		if i > len(left)-1 && j <= len(right)-1 {
			slice[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			slice[k] = left[i]
			i++
		} else if left[i] < right[j] {
			slice[k] = left[i]
			i++
		} else {
			slice[k] = right[j]
			j++
		}
	}
	fmt.Println("List 1: ", left)
	fmt.Println("List 2: ", right)
	fmt.Println("Merged: ", slice)
}
