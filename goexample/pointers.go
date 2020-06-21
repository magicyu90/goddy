package main

import "fmt"

// 指针
func zeroval(ival int) {
	ival = 0
}

// 指针
func zeroptr(iptr *int) {
	*iptr = 0
}

// 主函数
func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
}
