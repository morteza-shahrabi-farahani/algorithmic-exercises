package main

import "fmt"

func main() {
	var names []string
	var name string
	var peopleNumber int
	fmt.Scanf("%d", &peopleNumber)
	for i := 0; i < peopleNumber; i++ {
		fmt.Scanf("%s", &name)
		names = append(names, name)
	}

	for i := 0; i < peopleNumber - 1; i++ {
		fmt.Printf("%s to %s: ke ba in dar agar dar bande dar manand, dar manand.\n", names[i], names[i + 1])
		for j := i + 1; j > 0; j-- {
			fmt.Printf("%s to %s: dar manand?\n", names[j], names[j - 1])
		}
		for j := 0; j <= i; j++ {
			fmt.Printf("%s to %s: dar manand.\n", names[j], names[j + 1])
		}
	}
}