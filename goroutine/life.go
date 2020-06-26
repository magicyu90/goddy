package main

import (
	"fmt"
	"runtime"
)

// 一段时间的消费者
func consumer(ch chan int) {

	for {
		data := <-ch
		if data == 0 {
			break
		}

	}
	fmt.Println("goroutine exit")
}

func main() {
	// 创建通道
	ch := make(chan int)

	for {

		var input string
		fmt.Scan(&input)

		if input == "quit" {
			// 向所有的goroutine发送退出指令:0
			for i := 0; i < runtime.NumGoroutine()-1; i++ {
				ch <- 0
			}
			continue
		}
		go consumer(ch)

		fmt.Printf("当前goroutine数量:%d\n", runtime.NumGoroutine())

	}

}
