package main

import "fmt"

func main() {
	var inputNumber int
	var startCounter int
	fmt.Scanf("%d", &inputNumber)
	startCounter = (inputNumber / 2)
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

	for i := 0; i < inputNumber; i++ {
		fmt.Printf("*")
	}

	for i := 0; i < inputNumber; i++ {
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
