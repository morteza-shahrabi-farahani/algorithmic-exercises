package main

import (
	"fmt"
)

func main() {
	var number1, number2 int
	var list1 []string
	var list2 []string
	var counter1, counter2 int
	var holder string

	fmt.Scanf("%v %v", &number1, &number2)

	for i := 0; i < number1; i++ {
		fmt.Scanf("%s", &holder)
		for j := 0; j < len(holder); j++ {
			if holder[j] == '*' {
				counter1++
			}
		}
		list1 = append(list1, holder)
	}

	for i := 0; i < number1; i++ {
		fmt.Scanf("%s", &holder)
		for j := 0; j < len(holder); j++ {
			if holder[j] == '*' {
				counter2++
			}
		}
		list2 = append(list2, holder)
	}

	fmt.Printf("%v %v", counter1, counter2)
}