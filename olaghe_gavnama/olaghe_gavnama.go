package main

import "fmt"

func main() {
	var t, a, b int
	var ararCount, moomooCount int
	fmt.Scanf("%d %d %d", &t, &a, &b)
	sumCount := t / (1 + a + 1 + b)
	ararCount += sumCount
	moomooCount += sumCount
	reminder := t % (1 + a + 1 + b)
	if reminder > (1 + a) {
		ararCount++
		moomooCount++
	} else if reminder >= 1 {
		ararCount++
	}

	fmt.Println(ararCount, moomooCount)
}
