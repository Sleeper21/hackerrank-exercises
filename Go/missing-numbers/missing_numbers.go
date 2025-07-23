package main

import (
	"fmt"
)

/*
 * Complete the 'missingNumbers' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY arr
 *  2. INTEGER_ARRAY brr
 */

func missingNumbers(arr []int32, brr []int32) []int32 {
	// Write your code here
	numberCounter := make(map[int32]int32)

	for _, numB := range brr {
		numberCounter[numB]++
	}

	for _, numA := range arr {
		_, exists := numberCounter[numA]
		if exists {
			numberCounter[numA]--
		}

	}

	//  TODO: Continue

	return []int32{}
}

func main() {
	fmt.Println("Hello, missing numbers!")
}
