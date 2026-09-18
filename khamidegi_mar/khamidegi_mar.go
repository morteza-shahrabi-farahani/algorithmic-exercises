package main

import "fmt"

func main() {
	var snakeMap [2][8]int
	var inputStr string
	var row, column int
	var dead bool
	row = 1
	fmt.Scanf("%s", &inputStr)
	snakeMap[row][column] = 1
	for i := 0; i < len(inputStr); i++ {
		if inputStr[i] == 'R' {
			row += 1
			column += 1
			if row > 1 || column > 7 {
				fmt.Print("DEATH")
				dead = true
				break
			}
			snakeMap[row][column] = 1
		} else if inputStr[i] == 'L' {
			row -= 1
			column += 1
			if row < 0 || column > 7 {
				fmt.Print("DEATH")
				dead = true
				break
			}
			snakeMap[row][column] = 1
		} else if inputStr[i] == 'F' {
			column += 1
			snakeMap[row][column] = 1
			if row > 1 || column > 7 {
				fmt.Print("DEATH")
				dead = true
				break
			}
		}
	}

	if !dead {
		for i := 0; i < 2; i++ {
			for j := 0; j < 8; j++ {
				fmt.Print(snakeMap[i][j])
			}

			if i == 0 {
				fmt.Println()
			}
		}
	}
}
