package main

import "fmt"

func main() {
	var number1, number2, number3, number4, number5, number6, number7 int
	var numbers []int
	var chosenNumber int
	fmt.Scanf("%d %d %d %d %d %d %d\n", &number1, &number2, &number3, &number4, &number5, &number6, &number7)
	numbers = append(numbers, number1, number2, number3, number4, number5, number6, number7)
	fmt.Scanf("%d", &chosenNumber)

	for i := 0; i < len(numbers); i++ {
		if numbers[i] == chosenNumber {
			if i == 0 {
				fmt.Println("6")
				break
			}

			fmt.Println(7 - i)
			break
		}
	}
}
