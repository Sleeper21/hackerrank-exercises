package main

import (
	"fmt"
	"os"
)

func main() {
	// numCandles := 4
	// candles := []int32{3, 2, 1, 3}
	// Expected output: 2

	numCandles := 10
	candles := []int32{18, 90, 90, 13, 90, 75, 90, 8, 90, 43}
	// Expected output: 5

	if len(candles) != numCandles {
		os.Exit(1)
	}

	numOfBiggerCandle := birthdayCakeCandles(candles)

	fmt.Println(numOfBiggerCandle)
}

func birthdayCakeCandles(candles []int32) int32 {
	biggerCount := int32(0)
	var bigger int32

	for i, num := range candles {

		if i == 0 {
			bigger = num
			biggerCount++
			continue
		}

		if num > bigger {
			bigger = num
			biggerCount = 0
			biggerCount++
			continue
		}

		if num < bigger {
			continue
		}

		if num == bigger {
			biggerCount++
		}
	}

	return int32(biggerCount)
}
