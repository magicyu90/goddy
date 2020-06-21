package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

// 闭包
func closure() func(int) int {
	var x int
	// 匿名函数
	f := func(a int) int {
		x++
		return a + x
	}
	return f
}

// 返回加减函数
func calc(base int) (func(int) int, func(int) int) {

	fmt.Printf("%p\n", &base)

	// 加函数
	add := func(i int) int {
		fmt.Printf("%p\n", &base)
		base += i
		return base
	}

	// 减函数
	sub := func(i int) int {
		fmt.Printf("%p\n", &base)
		base -= i
		return base
	}
	return add, sub
}

//
func main() {
	// nextInt := intSeq()
	// fmt.Println(nextInt())
	// fmt.Println(nextInt())
	// newInts := intSeq()
	// fmt.Println(newInts())
	add, sub := calc(100)
	fmt.Println(add(1), sub(10))
	fmt.Println(add(3), sub(5))
	fmt.Println(add(7), sub(5))
	fmt.Println(add(10), sub(3))
}
