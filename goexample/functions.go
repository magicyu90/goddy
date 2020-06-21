package main

import "fmt"

// 相加函数
func plus(a int, b int) int {
	return a + b
}

func main() {
	res := plus(2, 3)
	fmt.Printf("res=%v\n", res)
}
