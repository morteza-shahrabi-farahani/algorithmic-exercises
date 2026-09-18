package main

import "fmt"

func main() {
	var inputNumber int
	result := 1
	fmt.Scanf("%d", &inputNumber)

	for i := 1; i <= inputNumber; i++ {
		result *= i
	}

	fmt.Println(result)
}
