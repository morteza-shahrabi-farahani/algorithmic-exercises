package main

import "fmt"

func main() {
	var start, end int
	fmt.Scanf("%d\n %d", &start, &end)
	if start == 1 {
		if end == 4 {
			fmt.Println(2)
		} else if end == 2 || end == 3 {
			fmt.Println(1)
		}
	}

	if start == 2 {
		if end == 3 {
			fmt.Println(2)
		} else if end == 1 || end == 4 {
			fmt.Println(1)
		}
	}

	if start == 3 {
		if end == 2 {
			fmt.Println(2)
		} else if end == 1 || end == 4 {
			fmt.Println(1)
		}
	}

	if start == 4 {
		if end == 1 {
			fmt.Println(2)
		} else if end == 2 || end == 3 {
			fmt.Println(1)
		}
	}

	if start == end {
		fmt.Println(0)
	}
}
