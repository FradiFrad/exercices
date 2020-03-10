package main

import "fmt"

// Complete the sockMerchant function below.
// Parameters :
//  - n: the number of socks in the pile (1 ≤ n ≤ 100)
//  - ar: the colors of each sock (1 ≤ arr[i] ≤ 100 where 0 ≤ i ≤ n)
// Returns an integer representing the number of matching pairs of socks available.
func sockMerchant(n int32, ar []int32) int32 {
	colorFrequency := make(map[int32]int32)
	pairs := int32(0)

	for _, color := range ar {
		fmt.Printf("------------ color : %d \n", color)
		_, ok := colorFrequency[color]
		if ok {
			pairs++
			delete(colorFrequency, color)
		} else {
			colorFrequency[color]  = 1
		}
	}
	
	fmt.Printf("pairs : %d \n", pairs)


	return pairs
}

func main() {
	ar := []int32{10, 20, 20, 10, 10, 30, 50, 10, 2}
	fmt.Println(sockMerchant(9, ar))
}
