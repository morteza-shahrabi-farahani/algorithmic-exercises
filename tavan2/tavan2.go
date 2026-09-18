package main

import "fmt"

func main() {
	var inputNumber int
	currentConditionNumber := 1
	fmt.Scanf("%d", &inputNumber)

	for i := 1; currentConditionNumber <= inputNumber; i++ {
		currentConditionNumber *= 2
	}

	fmt.Println(currentConditionNumber)
}
