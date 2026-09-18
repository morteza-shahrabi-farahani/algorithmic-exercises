package main

import "fmt"

func main() {
	var input1, input2 int
	fmt.Scanf("%d %d", &input1, &input2)
	if input1%input2 == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
