package main

import "fmt"

func main() {
	var inputString string
	fmt.Scanf("%s", &inputString)
	fmt.Printf("saal: %s\n", string(inputString[0])+string(inputString[1]))
	fmt.Printf("maah: %s\n", string(inputString[2])+string(inputString[3]))
}
