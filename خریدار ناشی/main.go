package main

import "fmt"


func main() {
	var row, column, notWorked int
	var tempX, tempY int
	var notWorkingX, notWorkingY []int
	var same, jump bool
	fmt.Scanf("%d %d %d", &row, &column, &notWorked)
	for i := 0; i <notWorked; i++ {
		fmt.Scanf("%d %d", &tempX, &tempY)
		notWorkingX = append(notWorkingX, tempX)
		notWorkingY = append(notWorkingY, tempY)
	}

	if notWorked % 2 != 0 && notWorked != row * column {
		fmt.Printf("0")
	} else if notWorked == row * column {
		if notWorked % 2 != 0 && row * column == notWorked {
			fmt.Printf("0")
		} else {
			fmt.Printf("-1")
		}
	} else {
		for i := 0; i <row; i++ {
			same = false
			for j := 0; j <column; j++ {
				same = false
				for z := 0; z < notWorked; z++ {
					if notWorkingX[z] == (i + 1) && notWorkingY[z] == (j + 1) {
						same = true
						break
					}
				}
				if !same {
					fmt.Printf("1\n")
					fmt.Printf("%d %d", (i + 1), (j + 1))
					jump = true
					break
				}
			}
			if jump {
				break
			}
		}
	}
}