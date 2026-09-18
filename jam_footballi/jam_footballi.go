package main

import "fmt"

func main() {
	var testCount int
	fmt.Scanf("%d\n", &testCount)
	for i := 0; i < testCount; i++ {
		var a, b, c, d int
		fmt.Scanf("%d %d %d %d\n", &a, &b, &c, &d)
		if a+c > b+d {
			fmt.Println("perspolis")
		} else if b+d > a+c {
			fmt.Println("esteghlal")
		} else if b > c {
			fmt.Println("esteghlal")
		} else if c > b {
			fmt.Println("perspolis")
		} else {
			fmt.Println("penalty")
		}
	}
}
