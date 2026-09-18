package main

import (
	"fmt"
	"strings"
)

func main() {
	var inputStr string
	fmt.Scanf("%s", &inputStr)
	result := strings.ReplaceAll(inputStr, ".NET", "Golang")
	fmt.Print(result)
}
