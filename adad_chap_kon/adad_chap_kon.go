package main

import (
	"fmt"
	"strconv"
)

func main() {
	var inputNumber string
	fmt.Scanf("%s", &inputNumber)

	for i := 0; i < len(inputNumber); i++ {
		currentNumber, _ := strconv.Atoi(string(inputNumber[i]))
		fmt.Printf("%d: ", currentNumber)
		for j := 0; j < currentNumber; j++ {
			fmt.Printf("%d", currentNumber)
		}

		fmt.Println()
	}
}
