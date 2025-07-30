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
	houseStart := int32(7)
	houseEnd := int32(11)
	appleTreePosition := int32(5)
	orangeTreePosition := int32(15)
	apples := []int32{-2, 2, 1}
	oranges := []int32{5, -6}

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
