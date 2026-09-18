package main

import "fmt"

func main() {
	var classesCounts int
	fmt.Scanf("%d\n", &classesCounts)
	for i := 0; i < classesCounts; i++ {
		var studentsCount int
		var names []string
		alphabetsCount := make(map[rune]int)
		var result int

		fmt.Scanf("%d\n", &studentsCount)

		for j := 0; j < studentsCount; j++ {
			var inputStr string
			innerAlphabetsCount := make(map[rune]int)

			fmt.Scanf("%s\n", &inputStr)
			names = append(names, inputStr)
			for _, char := range inputStr {
				innerAlphabetsCount[char]++
			}

			for key, _ := range innerAlphabetsCount {
				if innerAlphabetsCount[key] > alphabetsCount[key] {
					alphabetsCount[key] = innerAlphabetsCount[key]
				}
			}

			// fmt.Println(innerAlphabetsCount)
			// fmt.Println(alphabetsCount)
		}

		for _, value := range alphabetsCount {
			result += value
		}

		fmt.Println(result)
	}
}
