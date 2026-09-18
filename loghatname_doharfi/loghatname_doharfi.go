package main

import "fmt"

func main() {
	var input int
	fmt.Scanf("%d", &input)
	if input%2 == 0 {
		fmt.Println("b")
	} else {
		fmt.Println("a")
	}
}
