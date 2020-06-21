package main

import "fmt"

func vals() (int, int) {
	return 3, 7
}

// 主函数
func main() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// 如果只想需要返回值的一部分，使用空白符
	_, c := vals()
	fmt.Println(c)
}
