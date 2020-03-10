package main

import "fmt"
import "math"

// Complete the sockMerchant function below.
// Parameters :
//  - n: the number of socks in the pile (1 ≤ n ≤ 100)
//  - ar: the colors of each sock (1 ≤ arr[i] ≤ 100 where 0 ≤ i ≤ n)
// Returns an integer representing the number of matching pairs of socks available.
func sockMerchant(n int32, ar []int32) int32 {
	colorFrequency := make(map[int32]int32)
	for _, color := range ar {
		// Color is in the map
		_, ok := colorFrequency[color]
		if ok {
			colorFrequency[color] = colorFrequency[color] + 1
		} else {
			colorFrequency[color]  = 1
		}
	}
	fmt.Printf("colorFrequency : %d \n", colorFrequency)

	pairs := int32(0)
	for _, frequency := range colorFrequency {
		if frequency%2 != 0 {
			pairs += (frequency - 1)/2
		} else {
			pairs += frequency / 2
		}
	}
	fmt.Printf("pairs : %d \n", pairs)

	return 3
}

func isContainingColor(color int32, uniqueColors []int32) bool {
	for _, val := range uniqueColors {
		fmt.Printf("couleur  = %d // valeur = %d \n", color, val)
		if val != color {
			fmt.Printf("uniqueColors : %d ne contient pas encore la couleur %d (ici la valeur est : %d) \n ---------------------- ", uniqueColors, color, val)
			// fmt.Printf("value of uniqueColors %d is different from this color %d \n", val, color)
			return false
			// if !isContainingColor(color, uniqueColors) {
			// 	// arraypush equivalent
			// uniqueColors = append(uniqueColors, color)
			// }
		}
		fmt.Printf("uniqueColors : %d contient déjà la couleur %d \n ---------------------- ", uniqueColors, color)
		// }
	}
	return true
}

func countColors(color int32, socksColors []int32, uniqueColors []int32) []int32 {
	// var uniqueColors []int32
	// if len(uniqueColors) == 0 {
	// 	fmt.Printf("uniqueColors is empty ! \n")
	// 	uniqueColors = append(uniqueColors, color)
	// }

	// Si la couleur n'est pas déjà dans uniqueColors, on la met dans l'array uniqueColors

	for _, val := range uniqueColors {
		if val != color {
			fmt.Printf("value of uniqueColors %d is different from this color %d \n", val, color)

			// arraypush equivalent
			uniqueColors = append(uniqueColors, color)
		}
		// }
	}

	return uniqueColors
}

func countFrequency(color int32, socksColors []int32) int32 {
	var colorFrequency int32

	for _, value := range socksColors {
		if color == value {
			colorFrequency++
		}
	}
	return colorFrequency

}

func main() {
	ar := []int32{10, 20, 20, 10, 10, 30, 50, 10, 2}
	fmt.Println(sockMerchant(9, ar))
}
