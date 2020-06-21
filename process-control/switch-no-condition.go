package main

import (
	"time"
)

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		println("Good morning!")
	case t.Hour() < 17:
		println("Good afternoon!")
	default:
		println("Good evening!")
	}
}
