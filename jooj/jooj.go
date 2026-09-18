package main

import "fmt"

func main() {
	var numbersCount int
	var numbers []int
	var hasDastAndaz bool
	fmt.Scanf("%d\n", &numbersCount)
	for i := 0; i < numbersCount; i++ {
		var number int
		fmt.Scanf("%d", &number)
		numbers = append(numbers, number)
	}

	for i := 1; i < numbersCount-1; i++ {
		if numbers[i] > numbers[i-1] && numbers[i] > numbers[i+1] {
			fmt.Println("Ey baba :(")
			hasDastAndaz = true
			break
		}
	}

	if !hasDastAndaz {
		fmt.Println("Bah Bah! Ajab jooji!")
	}
}
