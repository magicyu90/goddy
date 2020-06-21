// 1 for 循环

package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Printf("sum和:%v", sum)
}
