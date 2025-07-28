package main

import (
	"fmt"
)

// Note:	- 12:00:00AM on a 12-hour clock is 00:00:00 on a 24-hour clock.
// 			- 12:00:00PM on a 12-hour clock is 12:00:00 on a 24-hour clock.

func main() {
	input := "07:05:45PM" // Expected output: 19:05:45

	output := convertTime(input)

	fmt.Println(output)
}

func convertTime(input string) string {
	var result, partialResult string
	var hourMapping map[string]string = make(map[string]string)
	hourMapping["01"] = "13"
	hourMapping["02"] = "14"
	hourMapping["03"] = "15"
	hourMapping["04"] = "16"
	hourMapping["05"] = "17"
	hourMapping["06"] = "18"
	hourMapping["07"] = "19"
	hourMapping["08"] = "20"
	hourMapping["09"] = "21"
	hourMapping["10"] = "22"
	hourMapping["11"] = "23"

	hourInput := input[:2]
	partOfDay := input[8:]
	partialResult = input[2:8]

	if hourInput == "12" && partOfDay == "AM" {
		result = fmt.Sprintf("00" + partialResult)

	} else if hourInput != "12" && partOfDay == "PM" {
		result = fmt.Sprintf(hourMapping[hourInput] + partialResult)

	} else {
		result = fmt.Sprintf(hourInput + partialResult)
	}

	return result
}
