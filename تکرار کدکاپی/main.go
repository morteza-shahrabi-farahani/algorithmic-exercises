package main

import "fmt"

func main() {
	var number int
	codeCup := "codecup6"

	fmt.Scanf("%d", &number)
	if number % 8 == 0 {
		fmt.Printf("6")
	} else {
		fmt.Printf("%c", codeCup[(number % 8) - 1])
	}
}