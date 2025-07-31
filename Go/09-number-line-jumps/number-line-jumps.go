package main

import (
	"fmt"
)

func main() {
	//input1 := []int{0, 2, 5, 3} expected output: "NO"
	//input2 := []int{0, 3, 4, 2} expected output: "YES"
	//input3 := []int{14 4 98 2} expected output: "YES"
	//input4 := []int{28 8 96 2} expected output: "NO"

	kang1 := int32(28)
	kang1Vel := int32(8)
	kang2 := int32(96)
	kang2Vel := int32(2)

	output := kangaroo(kang1, kang1Vel, kang2, kang2Vel)

	fmt.Println(output)
}

func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	// Write your code here

	for {

		if x1 == x2 {
			return "YES"
		}

		if x1 > x2 {
			x1 = x1 + v1
			x2 = x2 + v2
			if x2 > x1 {
				return "NO"
			}
		} else {
			x1 = x1 + v1
			x2 = x2 + v2
			if x1 > x2 {
				return "NO"
			}
		}
		continue
	}

	// -- SIMPLER EQUATION SOLUTION : -- //

	// x1 + n * v1 = x2 + n * v2
	//	=> x1 - x2 = n * (v2 - v1)
	// Ou seja
	// n = (x1 - x2) / (v2 - v1)

	// Para que o encontro aconteça:
	//	(x1 - x2) % (v2 - v1) == 0 → precisa ser divisível.
	//	n precisa ser um inteiro positivo ou zero (ou seja, ≥ 0).

	// if v1 == v2 {
	// 	// If velocities are the same they will only meet if they start in the same position
	// 	if x1 == x2 {
	// 		return "YES"
	// 	} else {
	// 		return "NO"
	// 	}
	// }

	//	if (x2 - x1) % (v1 - v2) == 0 && (x2 - x1) * (v1 - v2) < 0 {
	//		return "YES"
	//	}
	//	return "NO"
}
