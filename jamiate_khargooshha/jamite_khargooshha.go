package main

import "fmt"

func main() {
	var rabbitsCount, eatenRabbits, years int
	fmt.Scanf("%d %d\n %d", &rabbitsCount, &eatenRabbits, &years)
	newRabbits := rabbitsCount
	for i := 0; i < years; i++ {
		newRabbits = (newRabbits * 2) - eatenRabbits
	}

	fmt.Println(newRabbits)
}
