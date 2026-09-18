package main

import "fmt"

func main() {
	var inputNumber int
	var result int
	currentConditionNumber := 2
	fmt.Scanf("%d", &inputNumber)

	for i := 1; currentConditionNumber <= inputNumber; i++ {
		currentConditionNumber *= 2
		result++
	}

	fmt.Println(result)
}
