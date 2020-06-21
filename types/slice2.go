package main

import (
	"fmt"
)

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	// 创建新的切片，其长度为 2 个元素，容量为 4 个元素
	mySlice := slice1[1:3]
	fmt.Println(mySlice)
	// 使用原有的容量来分配一个新元素，将新元素赋值为 40
	mySlice = append(mySlice, 40)
	fmt.Println(mySlice)
}
