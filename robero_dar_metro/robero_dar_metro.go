package main

import "fmt"

func main() {
	var num1, num2, num3, num4, num5, num6, num7, num8 int
	var num9, num10, num11, num12, num13, num14, num15, num16 int
	var result int
	fmt.Scanf("%d %d %d %d %d %d %d %d\n %d %d %d %d %d %d %d %d", &num1, &num2, &num3, &num4, &num5, &num6, &num7, &num8,
		&num9, &num10, &num11, &num12, &num13, &num14, &num15, &num16)
	if num1 == 1 && num9 == 1 {
		result++
	}

	if num2 == 1 && num10 == 1 {
		result++
	}

	if num3 == 1 && num11 == 1 {
		result++
	}

	if num4 == 1 && num12 == 1 {
		result++
	}

	if num5 == 1 && num13 == 1 {
		result++
	}

	if num6 == 1 && num14 == 1 {
		result++
	}

	if num7 == 1 && num15 == 1 {
		result++
	}

	if num8 == 1 && num16 == 1 {
		result++
	}

	fmt.Println(result)
}
