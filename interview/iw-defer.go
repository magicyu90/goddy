package main

import (
	"fmt"
)

// func main() {
// 	defer_call()
// }

// 多个defer之间按照LIFO（后进先出）顺序进行执行
func defer_call() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()
	panic("触发异常")
}

// defer表达式中可以修改函数中的命名返回值
func c() (i int) {
	defer func() { i++ }()
	return 1
}

// 上面的示例程序，返回值变量名为i，在defer表达式中可以修改这个变量的值。
// 所以，虽然在return的时候给返回值赋值为1，后来defer修改了这个值，让i自增了1，
// 所以，函数的返回值是2而不是1。

func main() {
	fmt.Println(doubleScore(0))    //0
	fmt.Println(doubleScore(20.0)) //40
	fmt.Println(doubleScore(50.0)) //50
}

// 需要注意的是，函数的return value 不是原子操作.而是在编译器中分解为两部分：返回值赋值 和 return 。
// 而defer刚好被插入到末尾的return前执行。
// 故可以在derfer函数中修改返回值。如下示例：
func doubleScore(source float32) (score float32) {
	defer func() {
		if score < 1 || score >= 100 {
			//将影响返回值
			score = source
		}
	}()
	score = source * 2
	// defer在这里执行
	return
	//或者
	//return source * 2
}
