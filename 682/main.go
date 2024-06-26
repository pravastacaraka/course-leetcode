package main

import (
	"fmt"
	"strconv"
)

func calPoints(operations []string) int {
	var (
		result int
		temp   = make([]int, 0)
	)

	for _, v := range operations {
		switch v {
		case "C":
			temp = temp[:len(temp)-1]
		case "D":
			temp = append(temp, temp[len(temp)-1]*2)
		case "+":
			temp = append(temp, temp[len(temp)-1]+temp[len(temp)-2])
		default:
			d, _ := strconv.Atoi(v)
			temp = append(temp, d)
		}
	}

	for _, v := range temp {
		result += v
	}

	return result
}

func main() {
	fmt.Println(calPoints([]string{"1", "C"}))
}
