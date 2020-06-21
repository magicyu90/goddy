package main

import "fmt"

//求和方法
// 注意类型在变量名 之后。
func add(x, y int) int {
	return x + y
}

//main方法
func main() {
	fmt.Println(add(13, 42))
}
