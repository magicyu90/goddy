// 斐波纳契闭包

package main

import "fmt"

func fibonacci() func() int {
	x1, x2 := 0, 1
	sum := 0
	return func() int {
		sum = x1 + x2
		fmt.Printf("x1:%d,x2:%d,sum:%d\n", x1, x2, sum)
		x1 = x2
		x2 = sum
		return sum
	}
}

func main() {
	f := fibonacci() // 返回一个函数
	for i := 0; i < 10; i++ {
		fmt.Printf("第%d次:", i+1)
		f() // 执行函数
	}
}
