package main

import "fmt"

func main() {
	var inputNumber int
	var sum int
	fmt.Scanf("%d", &inputNumber)
	for i := 1; i < inputNumber; i++ {
		fmt.Printf("%d + ", i)
		sum += i
	}

	fmt.Printf("%d = ", inputNumber)
	sum += inputNumber
	fmt.Printf("%d", sum)
}
