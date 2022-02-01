package main

import "fmt"

func main() {
	var input int
	var ok bool
	ok = false
	fmt.Scanf("%d", &input)
	
	for input != 1 {
		if input % 2 == 0 {
			if input == 2 {
				ok = true
			} 
			input /= 2
		} else {
			ok = false
			break
		}
	}

	if input == 1 {
		ok = true
	}

	if ok == false {
		fmt.Printf("No")
	} else {
		fmt.Println("Yes")
	}
}