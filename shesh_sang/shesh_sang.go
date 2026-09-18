package main

import "fmt"

func main() {
	var inputStr string
	fmt.Scanf("%s", &inputStr)
	switch inputStr {
	case "space":
		fmt.Println("blue")
	case "mind":
		fmt.Println("yellow")
	case "reality":
		fmt.Println("red")
	case "power":
		fmt.Println("purple")
	case "time":
		fmt.Println("green")
	case "soul":
		fmt.Println("orange")
	}
}
