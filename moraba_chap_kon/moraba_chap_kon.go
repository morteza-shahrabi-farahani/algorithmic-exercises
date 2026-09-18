package main

import "fmt"

func main() {
	var inputNumber int
	fmt.Scanf("%d", &inputNumber)

	for i := 0; i < inputNumber; i++ {
		fmt.Printf("*")
	}

	fmt.Println()

	for i := 0; i < inputNumber-2; i++ {
		fmt.Printf("*")
		for j := 0; j < inputNumber-2; j++ {
			fmt.Printf(" ")
		}
		fmt.Printf("*\n")
	}

	for i := 0; i < inputNumber; i++ {
		fmt.Printf("*")
	}
}
