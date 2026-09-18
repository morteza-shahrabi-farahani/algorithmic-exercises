package main

import "fmt"

func main() {
	var inputNumber int
	fmt.Scanf("%d", &inputNumber)
	switch inputNumber {
	case 5:
		fmt.Println(2)
	case 2:
		fmt.Println(5)
	case 4:
		fmt.Println(3)
	case 3:
		fmt.Println(4)
	case 1:
		fmt.Println(6)
	case 6:
		fmt.Println(1)
	}
}
