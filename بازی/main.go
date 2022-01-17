package main

import "fmt"

func main() {
	var number int
	var numbersList []int
	var holder int

	fmt.Scanf("%d", &number)

	for i := 0; i < number; i++ {
		fmt.Scanf("%d", &holder)
		numbersList = append(numbersList, holder)
	}

	for i := 0; len(numbersList) > 0; i++ {
		numbersList = findMax(numbersList)
		if len(numbersList) == 0 {
			break
		}
		numbersList = findMin(numbersList)
	}
}

func findMax(numbersList []int) []int {
	var maximum = 0
	var maxIndex = 0

	for i := 0; i < len(numbersList); i++ {
		if maximum < numbersList[i] {
			maximum = numbersList[i]
			maxIndex = i
		}
	}

	fmt.Printf("%v ", maximum)

	return append(numbersList[:maxIndex], numbersList[maxIndex+1:]...)
}

func findMin(numbersList []int) []int {
	var minimum = 200
	var minIndex = 0

	for i := 0; i < len(numbersList); i++ {
		if minimum > numbersList[i] {
			minimum = numbersList[i]
			minIndex = i
		}
	}

	fmt.Printf("%v ", minimum)

	return append(numbersList[:minIndex], numbersList[minIndex+1:]...)
}