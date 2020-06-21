// nil切片
// 切片的零值是 nil。
// nil 切片的长度和容量为 0 且没有底层数组。
package main

import "fmt"

func main() {
	var s []int
	fmt.Println(s, len(s), cap(s)) //打印s值 s长度 和s容量
	if s == nil {
		fmt.Println("nil")
	}
}
