package main

import "fmt"


func main() {
	var length int
	var number int 
	var inputStr string
	fmt.Scanf("%d", &length)
	fmt.Scanf("%d", &number)
	fmt.Scanf("%s", &inputStr)

	//do function for number times
	for i := 0; i < number; i++ {
		//first add last element to first and shift other strings index
		inputStr = addLastToFirst(inputStr)
		inputStr = shiftRunes(inputStr)
	}

	fmt.Printf("%v", inputStr)
}

func addLastToFirst(input string) string {

	//cast string to byte slice for changing string
	castedString := []byte(input)
	lastElement := castedString[len(input) - 1]
	for i := len(input) - 1; i > 0; i-- {
		//shift from end of the string to first element of string
		castedString[i] = castedString[i - 1]
	}
	castedString[0] = lastElement

	return string(castedString)
}

func shiftRunes(input string) string {
	castedString := []byte(input)
	for i := 0; i < len(input); i++ {
		//shift from end of the string to first element of string
		if castedString[i] == 122 {
			castedString[i] = 97
		} else {
			castedString[i] = castedString[i] + 1
		}
	}

	return string(castedString)
}
