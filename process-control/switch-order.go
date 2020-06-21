package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("What's Saturday?")
	today := time.Now().Weekday()
	switch today {
	case time.Saturday:
		println("Today")
	case time.Friday:
		println("Friday")
	case time.Wednesday:
		println("Wednesday")
	case time.Thursday:
		println("Thursday")
	case time.Tuesday:
		println("Tuesday")
	case time.Monday:
		println("Monday")
	case time.Sunday:
		println("Sunday")
	default:
		println("Unknown")
	}
}
