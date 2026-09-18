package main

import "fmt"

func main() {
	var input string
	var count int
	fmt.Scanf("%s", &input)

	for _, char := range input {
		if char == 'a' || char == 'i' || char == 'o' || char == 'e' || char == 'u' {
			count++
		}
	}

	fmt.Println(count)
}
