// 向切片追加元素

package main

import "fmt"

func main() {
	var s []int
	printSlice(s) //切片

	s = append(s, 0) //添加一个空切片
	printSlice(s)

	s = append(s, 1) //这个切片按需增长
	printSlice(s)

	s = append(s, 2, 3, 4) //可以一次性添加多个元素
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
