package main

import "fmt"

func main() {
	for {
		var input int
		fmt.Scanf("%d\n", &input)
		if input == 0 {
			break
		}

		if input <= 1000000 {
			fmt.Println(input)
		} else if input > 1000000 && input <= 5000000 {
			fmt.Println(input - (input / 10))
		} else {
			fmt.Println(input - (input / 5))
		}
	}
}
