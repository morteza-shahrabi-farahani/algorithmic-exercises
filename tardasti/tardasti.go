package main

import "fmt"

func main() {
	var questionsCount int
	fmt.Scanf("%d\n", &questionsCount)
	for i := 0; i < questionsCount; i++ {
		var inputStr string
		var result int
		fmt.Scanf("%s\n", &inputStr)
		for j, char := range inputStr {
			if j == 0 && char == '0' {
				result++
			}

			if char == '0' && j > 0 && inputStr[j-1] == '1' {
				result++
			}
		}

		fmt.Println(result)
	}
}
