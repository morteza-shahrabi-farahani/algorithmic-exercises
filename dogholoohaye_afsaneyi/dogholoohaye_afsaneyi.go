package main

import (
	"fmt"
	"strings"
)

func reverse(s string) string {
	result := ""

	for i := len(s) - 1; i >= 0; i-- {
		result += string(s[i])
	}

	return result
}

func main() {
	var t int
	fmt.Scanf("%d", &t)

	for i := 0; i < t; i++ {
		var s, target string
		fmt.Scanf("%s %s", &s, &target)

		if len(s) != len(target) {
			fmt.Println("NO")
			continue
		}

		// حالت اول: target یک چرخش از s باشد
		if strings.Contains(s+s, target) {
			fmt.Println("YES")
			continue
		}

		// حالت دوم: target یک چرخش از reverse(s) باشد
		reversed := reverse(s)

		if strings.Contains(reversed+reversed, target) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
