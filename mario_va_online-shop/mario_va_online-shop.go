package main

import "fmt"

func main() {
	var coinCount, diamondCount int
	var requiredCoinsCount, requiredDiamondCount int
	var coinDiamondRate int

	fmt.Scanf("%d %d\n", &coinCount, &diamondCount)
	fmt.Scanf("%d %d\n", &requiredCoinsCount, &requiredDiamondCount)
	fmt.Scanf("%d\n", &coinDiamondRate)

	if ((diamondCount * coinDiamondRate) + coinCount) >=
		((coinDiamondRate * requiredDiamondCount) + requiredCoinsCount) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
