package main

import "fmt"

func main() {
	var find, input string
	var same = true
	var result bool
	fmt.Scanf("%s", &find)
	fmt.Scanf("%s", &input)
	for j := 0; j < len(find); j++ {
		same = true
		for i := 0; i < len(input); i++ {
			if find[(j + i) % (len(find))] == input[i] {
				continue
			} else {
				same = false
				break
			}
		}
		if same {
			result = true
		}
	}	

	if result {
		fmt.Printf("Yes")
	} else {
		fmt.Printf("No")
	}
}