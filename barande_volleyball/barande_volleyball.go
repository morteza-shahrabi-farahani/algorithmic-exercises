package main

import "fmt"

func main() {
	var setsCount int
	var temp int
	var setsPoints string
	points := make(map[rune]int)
	fmt.Scanf("%d\n", &setsCount)
	for i := 0; i < setsCount; i++ {
		fmt.Scanf("%d\n", &temp)
		fmt.Scanf("%s\n", &setsPoints)
		points['Q'] = 0
		points['C'] = 0
		for _, char := range setsPoints {
			points[char]++
		}

		if points['Q'] > points['C'] {
			fmt.Println("Quera")
		} else {
			fmt.Println("CodeCup")
		}
	}
}
