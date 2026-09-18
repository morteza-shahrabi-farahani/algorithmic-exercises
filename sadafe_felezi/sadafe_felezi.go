package main

import (
	"fmt"
	"math"
)

func main() {
	var t, w int
	var d float64
	var dominator, numerator float64
	fmt.Scanf("%d %d", &t, &w)

	numerator = float64(t * int(math.Pow(2, float64(w-1))))

	for ; w > 0; w-- {
		dominator += float64(math.Pow(2, float64(w-1)))
	}

	d = numerator / dominator
	fmt.Printf("%.4f\n", d)
}
