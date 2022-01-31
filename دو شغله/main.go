package main

import "fmt"

func main() {
	var day string
	fmt.Scanf("%s", &day)
	if day == "shanbe" || day == "doshanbe" || day == "chaharshanbe" {
		fmt.Printf("perspolis")
	} else if day == "yekshanbe" || day == "seshanbe" || day == "panjshanbe" {
		fmt.Printf("bahman")
	} else {
		fmt.Printf("tatil")
	}
}