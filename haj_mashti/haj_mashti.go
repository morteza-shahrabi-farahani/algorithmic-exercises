package main

import "fmt"

func main() {
	var inputStr string
	fmt.Scanf("%s", &inputStr)
	if inputStr[0] == 'Y' {
		fmt.Println("Haji")
	} else if inputStr[1] == 'Y' {
		fmt.Println("Karbalaee")
	} else if inputStr[2] == 'Y' {
		fmt.Println("Mashti")
	} else {
		fmt.Println("Agha")
	}
}
