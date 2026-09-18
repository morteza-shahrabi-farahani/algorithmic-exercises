package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanf("%d\n %d", &a, &b)
	var results []int
	for i := a + 1; i < b; i++ {
		if i == 1 {
			continue
		}

		var isPrime = true
		for j := 2; j <= i/2; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			results = append(results, i)
		}
	}

	for i := 0; i < len(results); i++ {
		fmt.Printf("%d", results[i])
		if i != len(results)-1 {
			fmt.Printf(",")
		}
	}
}
