package main

import "fmt"

func main() {
	var input int
	fmt.Scanf("%d", &input)
	if input%2 == 0 {
		fmt.Println("fard")
	} else if input == 3 || input == 5 || input == 7 || input == 11 || input == 13 || input == 17 || input == 19 || input == 23 || input == 29 || input == 31 ||
		input == 37 || input == 41 || input == 43 || input == 47 {
		fmt.Println("zoj")
	} else {
		fmt.Println("fard")
	}
}
