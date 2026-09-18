package main

import "fmt"

func main() {
	var questionsCount int
	var answers string
	var studentsCount int

	fmt.Scanf("%d\n", &questionsCount)
	fmt.Scanf("%s\n", &answers)
	fmt.Scanf("%d\n", &studentsCount)

	for i := 0; i < studentsCount; i++ {
		var studentResult int
		for j := 0; j < questionsCount; j++ {
			var answer string
			var sharpsCount int
			var chosenAnswer byte
			correctAnswer := answers[j]
			fmt.Scanf("%s\n", &answer)

			for z := 0; z < len(answer); z++ {
				if answer[z] == '#' {
					sharpsCount++
					switch z {
					case 0:
						chosenAnswer = 'A'
					case 1:
						chosenAnswer = 'B'
					case 2:
						chosenAnswer = 'C'
					case 3:
						chosenAnswer = 'D'
					}
				}
			}

			if sharpsCount > 1 {
				studentResult -= 1
			}

			if sharpsCount == 1 && chosenAnswer == correctAnswer {
				studentResult += 3
			}

			if sharpsCount == 1 && chosenAnswer != correctAnswer {
				studentResult -= 1
			}
		}

		fmt.Println(studentResult)
	}
}
