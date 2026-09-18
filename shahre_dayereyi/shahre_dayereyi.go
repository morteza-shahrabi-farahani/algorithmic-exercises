package main

import "fmt"

func main() {
	var n, m int
	fmt.Scanf("%d %d", &n, &m)

	hasInward := false
	hasOutward := false

	for i := 0; i < n; i++ {
		var direction int
		fmt.Scanf("%d", &direction)

		if direction == 0 {
			hasInward = true
		} else {
			hasOutward = true
		}
	}

	// جهت خیابان‌های دایره‌ای تأثیری در جواب ندارد.
	for i := 0; i < m; i++ {
		var direction int
		fmt.Scanf("%d", &direction)
	}

	if hasInward && hasOutward {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
