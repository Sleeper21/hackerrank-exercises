package main

import "fmt"

// Complete the 'diagonalDifference' function below.
// The function is expected to return an INTEGER.
// The function accepts 2D_INTEGER_ARRAY arr as parameter.

func main() {
	arr := [][]int32{
		{11, 2, 4},
		{4, 5, 6},
		{10, 8, -12},
	}

	fmt.Println("-- Result: ", diagonalDifference(arr))
}

func diagonalDifference(arr [][]int32) int32 {
	// Write your code here
	var dif int32
	//var diagonal2 []int32
	var diagonalSum1 int32
	var diagonalSum2 int32

	for i := range arr { // this is the same as: for i := 0; i < len(arr); i++
		diagonalSum1 = diagonalSum1 + arr[i][i]
		diagonalSum2 = diagonalSum2 + arr[i][(len(arr)-1)-i]
	}

	//diagonalSum1 = arr[0][0] + arr[1][1] + arr[2][2]
	//diagonalSum2 = arr[0][2] + arr[1][1] + arr[2][0]

	dif = diagonalSum1 - diagonalSum2
	if dif < 0 {
		dif = dif * -1
	}

	return dif

}
