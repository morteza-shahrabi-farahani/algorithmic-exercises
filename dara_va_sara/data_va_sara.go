package main

import "fmt"

func main() {
	var inputNumber int
	fmt.Scanf("%d\n", &inputNumber)
	if inputNumber == 1 {
		fmt.Println(2)
	} else {
		fmt.Println(3)
	}
}
