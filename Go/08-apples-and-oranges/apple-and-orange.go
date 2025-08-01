package main

import (
	"fmt"
	"os"
)

/*
 * Complete the 'countApplesAndOranges' function below.
 *
 * The function accepts following parameters:
 *  1. INTEGER s
 *  2. INTEGER t
 *  3. INTEGER a
 *  4. INTEGER b
 *  5. INTEGER_ARRAY apples
 *  6. INTEGER_ARRAY oranges
 */

func main() {
	houseStart := int32(5)
	houseEnd := int32(15)
	appleTreePosition := int32(3)
	orangeTreePosition := int32(20)
	apples := []int32{2, 3, 4, 5, 6, 7, 8, 9, 1, 0}
	oranges := []int32{-5, -6, -7, -8, -10}

	countApplesAndOranges(houseStart, houseEnd, appleTreePosition, orangeTreePosition, apples, oranges)

}

func countApplesAndOranges(houseStart int32, houseEnd int32, appleTreePosition int32, orangeTreePosition int32, apples []int32, oranges []int32) {

	if houseStart > houseEnd {
		err := fmt.Errorf("error: starting point must be lower than the end point")
		fmt.Println(err)
		os.Exit(1)
	}

	finalApplesPosition := getFallenFruitPosition(appleTreePosition, apples)
	finalOrangesPosition := getFallenFruitPosition(orangeTreePosition, oranges)

	applesInHouse := countFruitsInHouse(finalApplesPosition, houseStart, houseEnd)
	orangesInHouse := countFruitsInHouse(finalOrangesPosition, houseStart, houseEnd)

	fmt.Println(applesInHouse)
	fmt.Println(orangesInHouse)

	// Print Christmas tree with base width equal to applesInHouse
	printChristmasTree(int(applesInHouse))
}

func getFallenFruitPosition(tree int32, coordinates []int32) []int32 {
	finalPositions := make([]int32, len(coordinates))

	for i, coordinate := range coordinates {
		finalPositions[i] = tree + coordinate
	}

	return finalPositions
}

func countFruitsInHouse(coordinates []int32, start int32, end int32) int32 {
	var count int32

	for _, coordinate := range coordinates {
		if coordinate >= start && coordinate <= end {
			count++
		} else {
			continue
		}
	}

	return count
}

// printChristmasTree prints a Christmas tree with '*' characters
// The base width equals the baseWidth parameter
func printChristmasTree(baseWidth int) {
	if baseWidth <= 0 {
		fmt.Println("No tree to print (0 apples in house)")
		return
	}

	fmt.Println("\nChristmas Tree:")

	// Calculate tree height based on base width
	height := baseWidth
	if baseWidth > 10 {
		height = 10 // Cap height for very large numbers
	}

	// Print the tree layers
	for i := 1; i <= height; i++ {
		spaces := height - i
		stars := 2*i - 1

		// Print spaces for centering
		for j := 0; j < spaces; j++ {
			fmt.Print(" ")
		}

		// Print stars
		for j := 0; j < stars; j++ {
			fmt.Print("*")
		}

		fmt.Println()
	}

	// Print the trunk (base)
	trunkWidth := baseWidth
	trunkSpaces := height - trunkWidth/2

	for j := 0; j < trunkSpaces; j++ {
		fmt.Print(" ")
	}

	for j := 0; j < trunkWidth; j++ {
		fmt.Print("*")
	}

	fmt.Println()
}
