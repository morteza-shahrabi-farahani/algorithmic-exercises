package main

import "fmt"

func main() {
	var inputNumber int
	fmt.Scanf("%d", &inputNumber)
	for i := 0; i < inputNumber; i++ {
		fmt.Printf("#")
	}
}
