package main

import "fmt"

func main() {
	var changesCount int
	var startingPosition int = 1
	fmt.Scanf("%d\n", &changesCount)
	for i := 0; i < changesCount; i++ {
		var start, end int
		fmt.Scanf("%d %d\n", &start, &end)
		if start == startingPosition {
			startingPosition = end
			continue
		}

		if end == startingPosition {
			startingPosition = start
			continue
		}
	}

	fmt.Println(startingPosition)
}
