package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scanf("%d\n %d\n %d", &a, &b, &c)
	if ((a*a + b*b) == c*c) || ((b*b + c*c) == a*a) || ((c*c + a*a) == b*b) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
