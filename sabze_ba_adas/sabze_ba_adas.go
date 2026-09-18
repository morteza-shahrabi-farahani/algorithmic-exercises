package main

import "fmt"

func main() {
	var n, k int
	fmt.Scanf("%d %d", &n, &k)
	if n%2 == 0 {
		if k <= n/2 {
			fmt.Println("0")
		} else {
			result := 1
			for i := (n / 2) + 2; i <= k; i++ {
				result += 2
			}

			fmt.Println(result)
		}
	} else {
		if k <= (n/2)+1 {
			fmt.Println("0")
		} else {
			result := 0
			for i := (n / 2) + 2; i <= k; i++ {
				result += 2
			}

			fmt.Println(result)
		}
	}
}
