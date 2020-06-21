//切片就像数组的引用
//更改切片的元素会修改其底层数组中对应的元素。

package main

import "fmt"

func main() {
	names := [4]string{
		"John", "Paul", "Kobe", "Renne",
	}

	fmt.Println(names)
	a := names[0:2]
	b := names[0:3]
	fmt.Println(a, b)

	b[0] = "James" // 改变当前数组第一个元素内容,改变切片的同时会影响原始数组
	fmt.Println(a, b)
	fmt.Println(names)
}
