package main

import "fmt"

func judgeCircle(moves string) bool {
	var pos = []int{0, 0} // refer x, y coordinate
	for _, c := range moves {
		switch c {
		case 'U':
			pos[1] += 1
		case 'D':
			pos[1] -= 1
		case 'R':
			pos[0] += 1
		case 'L':
			pos[0] -= 1
		}
	}
	return pos[0] == 0 && pos[1] == 0
}

func main() {
	fmt.Println(judgeCircle("LL"))
}
