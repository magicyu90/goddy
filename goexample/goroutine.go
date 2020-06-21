//协程
package main

import "fmt"

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

// 主函数
func main() {
	f("Hugo")

	// 协程
	go f("goroutine")

	// 匿名协程
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	var input string
	fmt.Scanln(&input)

	fmt.Println("done")
}
