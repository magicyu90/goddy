package main

import "fmt"

func do(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("这是int类型")
	case string:
		fmt.Println("这是string类型")
	default:
		fmt.Println("未知类型")
	}
}

func main() {
	do(21)
	do("hello")
	do(float64(2))
}
