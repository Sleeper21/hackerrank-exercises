package main

import (
	"fmt"
	"strings"
)

func main() {
	staircase(6)
}

//Complete the 'staircase' function below.
//The function accepts INTEGER n as parameter.

func staircase(n int32) {
	var spaces, hashes string

	for i := n - 1; i >= 0; i-- {

		if i == 0 {
			spaces = strings.Repeat(" ", int(i))
			hashes = strings.Repeat("#", int(n-i))

		} else {
			spaces = strings.Repeat(" ", int(i))
			hashes = strings.Repeat("#", int(n-i))
		}

		fmt.Println(spaces + hashes)
	}
}
