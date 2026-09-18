package main

import "fmt"

func main() {
	var inputNumber string
	var result = true
	fmt.Scanf("%s", &inputNumber)
	for i := 0; i < len(inputNumber)/2; i++ {
		if inputNumber[i] != inputNumber[len(inputNumber)-i-1] {
			result = false
		}
	}

	if result {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
