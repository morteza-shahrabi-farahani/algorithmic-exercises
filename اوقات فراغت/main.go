package main

import "fmt"

func main() {
	var row, column int
	var word string
	var words []string
	var find string
	var result int
	var same bool
	fmt.Scanf("%d %d", &row, &column)
	for i := 0; i < row; i++ {
		fmt.Scanf("%s", &word)
		words = append(words, word)
	}

	fmt.Scanf("%s", &find)
	// fmt.Printf("%s", find)
	// for i := 0; i < row; i++ {
	// 	fmt.Printf("%s", words[i])
	// }
	
	for i := 0; i < row; i++ {
		// fmt.Printf("hello\n")
		// fmt.Printf("len is %d\n", len(words[i]))
		for j := 0; j <= len(words[i]) - len(find); j++ {
			// fmt.Printf("hi\n")
			same = true
			for z := 0; z < len(find); z++ {
				// fmt.Printf("first char is: %c and second char is %c\n", words[i][z + j], find[z])
				if words[i][z + j] != find[z] {
					same = false
					break 
				}
			}
			if same == true {
				result++
			}
		}
	}

	for i := 0; i < column; i++ {
		for j := 0; j <= row - len(find); j++ {
			same = true
			for z := 0; z < len(find); z++ {
				if words[j + z][i] != find[z] {
					same = false
					break
				}
			}
			if same == true {
				result++
			}
		}
	}

	fmt.Printf("%v ", result)
}