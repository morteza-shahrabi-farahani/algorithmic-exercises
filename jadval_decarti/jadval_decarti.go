package main

import "fmt"

func main() {
	var row, column int
	fmt.Scanf("%d %d", &row, &column)
	for j := 0; j < row; j++ {
		for i := 0; i < column; i++ {
			fmt.Print(" _")
		}
		fmt.Println(" ")

		for i := 0; i < column; i++ {
			fmt.Print("| ")
		}
		fmt.Println("|")
	}

	for i := 0; i < column; i++ {
		fmt.Print(" _")
	}
	fmt.Println(" ")
}
