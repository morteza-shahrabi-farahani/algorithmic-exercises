package main

import "fmt"

func main() {
	var sitesCount, sitesAttendeesCount, siteAttendeesCost int
	var finalCost int64
	fmt.Scanf("%d\n", &sitesCount)
	for i := 0; i < sitesCount; i++ {
		fmt.Scanf("%d %d\n", &sitesAttendeesCount, &siteAttendeesCost)
		finalCost += int64(sitesAttendeesCount * siteAttendeesCost)
	}

	fmt.Println(finalCost)
}
