// 切片的长度与容量

package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	s = s[:0] //截取切片使其长度为0
	printSlice(s)

	s = s[:4] //拓展其长度
	printSlice(s)

	s = s[2:] //舍弃前两个值
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("长度=%d 容量=%d %v\n", len(s), cap(s), s)
}
