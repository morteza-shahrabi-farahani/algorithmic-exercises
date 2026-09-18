package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)
	if a == 0 && b == 0 {
		fmt.Println("infinite")
	} else if a == 0 {
		fmt.Println("invalid")
	} else {
		fmt.Println("unique")
	}
}
