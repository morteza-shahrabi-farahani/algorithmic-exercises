package main

import "fmt"

func main() {
	var testsCount int
	fmt.Scanf("%d\n", &testsCount)
	for i := 0; i < testsCount; i++ {
		var a, b, h int
		totalDays := 1
		fmt.Scanf("%d %d %d\n", &a, &b, &h)
		height := 0
		for j := 0; height < h; j++ {
			height += a
			if height >= h {
				break
			}

			height -= b
			totalDays++
		}

		fmt.Println(totalDays)
	}
}
