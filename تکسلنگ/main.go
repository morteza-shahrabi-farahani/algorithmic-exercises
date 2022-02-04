package main

import "fmt"

func main() {
	var first, second, third string
	fmt.Scanf("%s", &first)
	fmt.Scanf("%s", &second)
	fmt.Scanf("%s", &third)

	for i := 0; i < len(first) / 5; i++ {
		// fmt.Printf("hello\n")
		// fmt.Printf("%d\n", len(first))
		if first[i * 5] == '*' && first[(i * 5) + 1] == '*' && first[(i * 5) + 2] == '*' && first[(i * 5) + 3] == '*' && first[(i * 5) + 4] == '*' {
			fmt.Printf("T")
			continue
		} else if first[i * 5] == '*' && first[(i * 5) + 1] == '*' && first[(i * 5) + 2] == 'o' && first[(i * 5) + 3] == '*' && first[(i * 5) + 4] == '*' {
			fmt.Printf("M")
			continue
		} else if first[i * 5] == 'o' && first[(i * 5) + 1] == 'o' && first[(i * 5) + 2] == '*' && first[(i * 5) + 3] == 'o' && first[(i * 5) + 4] == 'o' {
			fmt.Printf("A")
			continue
		} else {
			if second[i*5] == 'o' {
				fmt.Printf("X")
				continue
			} else {
				fmt.Printf("N")
				continue
			}
		}
	}
}