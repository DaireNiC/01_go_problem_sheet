//Author: Daire Ní Chatháin
//Problem sheet 1 : Largest and Smallest in Int

/* This Program that returns the largest  and smallest elements in a list.

Resource: https://golang.org/pkg/sort/

*/

package main

import "fmt"
import "sort"

func main() {
	
	//an unsorted array of integers
	s := []int{5, 2, 6, 3, 1, 4}
	//print unsorted values
	fmt.Println("Unsorted : ", s)
	
	//sort using inbuilt sort method for ints
	sort.Ints(s)
	
	//print sorted result to console 
	fmt.Println("Sorted: ", s)
}
