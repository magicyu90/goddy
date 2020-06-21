// 变参函数

package main

import "fmt"

// 可变参数
func sum(nums ...int) {
	fmt.Println(nums, "")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

// 主函数
func main() {
	sum(1)
	sum(1, 2)
	sum(1, 2, 3, 4)
}
