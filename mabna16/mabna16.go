package main

import (
	"fmt"
)

func main() {
	var inputStr string
	fmt.Scanf("%s", &inputStr)

	hexChars := "0123456789ABCDEF"
	hexMap := make(map[byte]int)
	for i := 0; i < len(hexChars); i++ {
		hexMap[hexChars[i]] = i
	}

	// تبدیل رشته به یک آرایه از بایت‌ها تا بتوانیم از انتها افزایش دهیم
	bytes := []byte(inputStr)
	carry := 1

	for i := len(bytes) - 1; i >= 0; i-- {
		v := hexMap[bytes[i]] + carry
		if v == 16 {
			bytes[i] = '0'
			carry = 1
		} else {
			bytes[i] = hexChars[v]
			carry = 0
			break
		}
	}

	// اگر هنوز حمل مانده بود، یک '1' در ابتدای نتیجه اضافه می‌کنیم
	if carry == 1 {
		bytes = append([]byte{'1'}, bytes...)
	}

	fmt.Println(string(bytes))
}
