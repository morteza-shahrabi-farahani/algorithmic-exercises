package main

import (
	"fmt"
	"strings"
)

func isValid(s string) bool {
	// حالت اول: 09xxxxxxxxx
	if strings.HasPrefix(s, "09") {
		if len(s) != 11 {
			return false
		}

		// از آنجایی که ممکن است + وسط رشته باشد،
		// بررسی می‌کنیم همه‌ی کاراکترها رقم باشند.
		for _, c := range s {
			if c < '0' || c > '9' {
				return false
			}
		}

		return true
	}

	// حالت دوم: 98xxxxxxxxx
	if strings.HasPrefix(s, "98") {
		if len(s) != 12 {
			return false
		}

		for _, c := range s {
			if c < '0' || c > '9' {
				return false
			}
		}

		return true
	}

	// حالت سوم: +98xxxxxxxxx
	if strings.HasPrefix(s, "+98") {
		// + یک کاراکتر است، بنابراین طول رشته باید 12 باشد
		if len(s) != 13 {
			return false
		}

		// بعد از + باید همه‌ی کاراکترها رقم باشند
		for i, c := range s {
			if i == 0 {
				continue
			}

			if c < '0' || c > '9' {
				return false
			}
		}

		return true
	}

	return false
}

func convert(s string) string {
	if strings.HasPrefix(s, "09") {
		// 09123456789
		// ↓
		// +989123456789
		return "+98" + s[1:]
	}

	if strings.HasPrefix(s, "98") {
		// 989123456789
		// ↓
		// +989123456789
		return "+" + s
	}

	// قبلاً معتبر بودنش بررسی شده است
	return s
}

func main() {
	var n int
	var phoneNumber string
	fmt.Scanf("%d\n", &n)

	for i := 0; i < n; i++ {
		fmt.Scanf("%s\n", &phoneNumber)

		if !isValid(phoneNumber) {
			fmt.Println("invalid")
		} else {
			fmt.Println(convert(phoneNumber))
		}
	}
}
