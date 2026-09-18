package main

import "fmt"

func main() {
	var input int
	fmt.Scanf("%d", &input)
	if input == 1 || input == 2 || input == 3 || input == 4 || input == 6 || input == 12 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
