package main

import (
	"fmt"
)

func isRobotBounded(instructions string) bool {
	var direction = "N"
	var pos = []int{0, 0} // represents x, y coordinate

	iter := 0
	for iter < 4 {
		for _, c := range instructions {
			switch c {
			case 'G':
				switch direction {
				case "N":
					pos[1] += 1
				case "S":
					pos[1] -= 1
				case "E":
					pos[0] += 1
				case "W":
					pos[0] -= 1
				}
			case 'L':
				switch direction {
				case "N":
					direction = "W"
				case "S":
					direction = "E"
				case "E":
					direction = "N"
				case "W":
					direction = "S"
				}
			case 'R':
				switch direction {
				case "N":
					direction = "E"
				case "S":
					direction = "W"
				case "E":
					direction = "S"
				case "W":
					direction = "N"
				}
			}
		}
		if pos[0] == 0 && pos[1] == 0 {
			return true
		}
		iter++
	}

	return pos[0] == 0 && pos[1] == 0
}

func main() {
	fmt.Println(isRobotBounded("GL"))
}
