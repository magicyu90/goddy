// go 并发模型 轻松入门流水线模型

package main

import (
	"fmt"
	"time"
)

// 生产者消费者模型
// 返回通道
func producer(nums ...int) <-chan int {
	// 创建通道
	out := make(chan int)
	go func() {
		defer close(out)
		for _, num := range nums {
			out <- num
		}
	}()
	return out
}

// 从通道里取 计算平方 并返回新通道
func square(inCh <-chan int) <-chan int {
	// 创建通道
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range inCh {
			out <- n * n
		}
	}()
	return out
}

// 主函数
func main() {
	// in:= producer(1,2,3,4)
	// ch:= square(in)

	// for n:= range ch{
	// 	fmt.Printf("%3d",n)
	// }
	// fmt.Println()

	channel2 := make(chan int, 0)
	go func() {
		time.Sleep(time.Second * 5)
		v := <-channel2
		fmt.Printf("the value is %v\n", v)
	}()
	channel2 <- 1
	fmt.Print("the time is over\n")
}
