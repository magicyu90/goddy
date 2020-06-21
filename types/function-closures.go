// 函数的闭包

package main

import "fmt"

func addr() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {

	fmt.Println("函数闭包")
}
