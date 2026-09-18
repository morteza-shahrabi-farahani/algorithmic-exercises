package main

import "fmt"

func main() {
	var numbersCount int
	var currentNumber int
	maximum := 0
	fmt.Scanf("%d", &numbersCount)
	for i := 0; i <= numbersCount; i++ {
		fmt.Scanf("%d", &currentNumber)
		if maximum < currentNumber {
			maximum = currentNumber
		}
	}

	fmt.Printf("%d", maximum)
}
