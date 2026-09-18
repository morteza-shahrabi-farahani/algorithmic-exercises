package main

import "fmt"

func main() {
	var numberOfQuestions int
	fmt.Scanf("%d\n", &numberOfQuestions)
	for i := 0; i < numberOfQuestions; i++ {
		var peopleCount, totalCost, governmentCost int
		fmt.Scanf("%d %d %d\n", &peopleCount, &totalCost, &governmentCost)
		if (totalCost+((peopleCount-1)*governmentCost))%peopleCount == 0 {
			if ((totalCost+((peopleCount-1)*governmentCost))/peopleCount)-governmentCost > 0 {
				fmt.Println(((totalCost + ((peopleCount - 1) * governmentCost)) / peopleCount) - governmentCost)
			} else {
				fmt.Println("-1")
			}
		} else {
			fmt.Println("-1")
		}
	}
}
