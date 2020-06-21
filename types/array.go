// 数组

package main

import "fmt"

func main() {
	var a [2]string //定义字符串数组
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 4, 11, 22, 37} //定义长度为6的整形数组
	fmt.Println(primes)
}
