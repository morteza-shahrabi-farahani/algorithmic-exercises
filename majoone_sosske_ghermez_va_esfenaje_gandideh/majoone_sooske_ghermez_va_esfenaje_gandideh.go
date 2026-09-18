package main

import "fmt"

func main() {
	var a, b, c, d, m int
	fmt.Scanf("%d %d %d %d %d", &a, &b, &c, &d, &m)
	if (a + (m * c)) > (b + (m * d)) {
		if c > d {
			fmt.Println("Eyval baba!")
		} else {
			fmt.Println("Naaa, eshtebahe!")
		}
	} else {
		if d > c {
			fmt.Println("Eyval baba!")
		} else {
			fmt.Println("Naaa, eshtebahe!")
		}
	}
}
