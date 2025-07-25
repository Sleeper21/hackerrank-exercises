package main

import "fmt"

func main() {
	input := []int32{1, 2, 3, 4, 5}

	miniMaxSum([5]int32(input))
}

// Complete the 'miniMaxSum' function below.

// The function accepts INTEGER_ARRAY arr as parameter with 5 positive integers.

func miniMaxSum(arr [5]int32) {
	// Write code here
	var arrToSum []int64
	var allSums [len(arr)]int64
	var minSum, maxSum int64

	for j := 0; j < len(arr); j++ {

		arrToSum = []int64{}

		for i, num := range arr {

			if i == j {
				continue
			}
			arrToSum = append(arrToSum, int64(num))
		}

		allSums[j] = sumNumbers(arrToSum)
	}

	minSum = findMin(allSums)
	maxSum = findMax(allSums)

	fmt.Printf("%v %v\n", minSum, maxSum)
}

func sumNumbers(nums []int64) int64 {
	var r int64

	for _, num := range nums {
		r += num
	}
	return r
}

func findMin(nums [5]int64) int64 {
	var minimum int64

	for i, num := range nums {
		if i == 0 {
			minimum = num
		} else {
			if num < minimum {
				minimum = num
			}
		}
	}
	return minimum
}

func findMax(nums [5]int64) int64 {
	var maximum int64

	for i, num := range nums {
		if i == 0 {
			maximum = num
		} else {
			if num > maximum {
				maximum = num
			}
		}
	}
	return maximum
}
