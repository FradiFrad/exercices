package main

import "fmt"

// Given :
// a list of numbers => slice of int
// and a number k => int
// return whether any two numbers from the list add up to k.

// For example, given [10, 15, 3, 7] and k of 17, return true since 10 + 7 is 17.
// Recursive !
func isItAddedUp(k int, numbers []int) bool {
	for _, number := range numbers {
		number1 := number
		numbers2 := append(numbers[:1], numbers[2:]...)
		for _, number2 := range numbers2 {
			if number1+number2 == k {
				return true
			}
		}
		// end of this loop = pas de somme égale à k

	}
	return false
}

func main() {
	fmt.Println(isItAddedUp(17, []int{10, 1}))
}
