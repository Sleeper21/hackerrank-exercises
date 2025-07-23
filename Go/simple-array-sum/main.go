package main

import (
	"fmt"
)

func main() {
	ar := []int32{1, 2, 3, 4, 10, 11}

	fmt.Println("-- Result: ", simpleArraySum(ar))
}

// Complete the 'simpleArraySum' function below.
//
// The function is expected to return an INTEGER.
// The function accepts INTEGER_ARRAY ar as parameter.

func simpleArraySum(ar []int32) int32 {
	// Write your code here
	var sum int32

	for _, num := range ar {
		sum = sum + num
	}

	return sum
}
