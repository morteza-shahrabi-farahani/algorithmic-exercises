package main

import "fmt"

func main() {
	var inputNumber int
	resultsSum := 0
	fmt.Scanf("%d", &inputNumber)
	for i := 1; i <= inputNumber/2; i++ {
		if inputNumber%i == 0 {
			resultsSum += i
		}
	}

	if resultsSum == inputNumber {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
