package main

import "fmt"

func main() {
	var number int
	var numbers []int
	var counter int
	var result bool
	fmt.Scanf("%d", &counter)
	for i := 0; i < counter; i++ {
		fmt.Scanf("%d", &number)
		numbers = append(numbers, number)
	}

	for i := 1; i < counter - 1; i++ {
		if numbers[i] > numbers[i - 1] && numbers[i] > numbers[i + 1]{
			result = true
		}
	}

	if result {
		fmt.Println("Ey baba :(")
	} else {
		fmt.Println("Bah Bah! Ajab jooji!")
	}
}