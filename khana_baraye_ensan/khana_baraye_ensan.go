package main

import "fmt"

func main() {
	var questionsCount int
	fmt.Scanf("%d\n", &questionsCount)
	for i := 0; i < questionsCount; i++ {
		var currentNumber int
		fmt.Scanf("%d\n", &currentNumber)
		if currentNumber < 1024 {
			fmt.Printf("%dB\n", currentNumber)
			continue
		}

		currentNumber /= 1024
		if currentNumber < 1024 {
			fmt.Printf("%dKiB\n", currentNumber)
			continue
		}

		currentNumber /= 1024
		if currentNumber < 1024 {
			fmt.Printf("%dMiB\n", currentNumber)
			continue
		}
	}
}
