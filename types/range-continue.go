// range（续）
// 可以将下标或值赋予 _ 来忽略它。
// for i, _ := range pow
// for _, value := range pow

package main

import "fmt"

func main() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i)
	}
	for index, value := range pow {
		fmt.Printf("索引:%d,值:%d\n", index, value)
	}
}
