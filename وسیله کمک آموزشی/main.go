package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var number int
	var line1, line2 string
	var lamps []string
	var status []string
	var answers []string
	var result []int

	fmt.Scanf("%d", &number)
	// fmt.Scanf("%s", &line1)
	// fmt.Scanf("%s", &line2)
	reader := bufio.NewReader(os.Stdin)
	line1, _ = reader.ReadString('\n')
	line2, _ = reader.ReadString('\n')
	// for scanner.Scan() {
	// 	line1 = scanner.Text()
	// 	line2 = scanner.Text()
	// }
	// temp := line1[len(line1) - 2]
	// fmt.Printf(string(temp))
	lamps = strings.Split(line1, " ")
	status = strings.Split(line2, " ")
	// fmt.Printf("%d\n", len(lamps))
	// fmt.Printf("%s\n", lamps[4])
	// fmt.Printf("%s\n", status[4])
	for i := 0; i < len(status); i++ {
		// fmt.Printf("%s %s\n", status[i], lamps[i])
		if strings.Contains(status[i], "1") {
			// fmt.Printf("hi\n")
			// fmt.Printf("%d\n", i)
			if i != len(status) - 1 {
				answers = append(answers, lamps[i])
			} else {
				// answers = append(answers, string(line1[len(line1) - 2]))
				temp := lamps[i][0: len(lamps[i]) - 1]
				answers = append(answers, temp)
			}
		}
		
	}

	// if line2[len(line1) - 1] == '1' {
	// 	answers = append(answers, string(line2[len(line2) - 1]))
	// }

	// sort.Strings(answers)
	// fmt.Printf("%d", len(answers))
	for i := 0; i < len(answers); i++ {
		j, err := strconv.Atoi(answers[i])
		if err != nil {
			// fmt.Printf("error")
			// fmt.Printf("%s\n", answers[len(answers) - 1])
		}

		result = append(result, j)
	}

	sort.Ints(result)

	for i := 0; i < len(result); i++ {
		fmt.Printf("%d ", result[i])
	}
}