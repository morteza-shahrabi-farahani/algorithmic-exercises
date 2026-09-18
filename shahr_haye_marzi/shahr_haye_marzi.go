package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanf("%d\n %d", &a, &b)
	if a == 1 {
		fmt.Println(b)
	} else if b == 1 {
		fmt.Println(a)
	} else {
		fmt.Println((2 * (a + b)) - 4)
	}
}
