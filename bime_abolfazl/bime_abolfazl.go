package main

import "fmt"

func main() {
	var inputStr string
	result := true
	fmt.Scanf("%s", &inputStr)
	for i := range inputStr {
		if inputStr[i] == 'm' {
			result = false
		}
	}

	if result {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
