package main

import "fmt"

func main() {
	var inputNumber int
	var startCounter int
	fmt.Scanf("%d", &inputNumber)
	newInputNumber := (2 * inputNumber) + 1
	startCounter = (newInputNumber / 2)
	for i := 0; i < startCounter; i++ {
		for j := 0; j < startCounter-i; j++ {
			fmt.Printf(" ")
		}

		for j := 0; j <= i; j++ {
			fmt.Printf("*")
		}

		for j := 0; j < i; j++ {
			fmt.Printf("*")
		}

		for j := 0; j < startCounter-i; j++ {
			fmt.Printf(" ")
		}

		fmt.Println()
	}

	for i := 0; i < newInputNumber; i++ {
		fmt.Printf("*")
	}

	fmt.Println()

	for i := startCounter - 1; i >= 0; i-- {
		for j := 0; j < startCounter-i; j++ {
			fmt.Printf(" ")
		}

		for j := 0; j <= i; j++ {
			fmt.Printf("*")
		}

		for j := 0; j < i; j++ {
			fmt.Printf("*")
		}

		for j := 0; j < startCounter-i; j++ {
			fmt.Printf(" ")
		}

		if i != 0 {
			fmt.Println()
		}
	}
}
