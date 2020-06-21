package main

import (
	"fmt"
)

// 主函数
func main() {

	// 消息通道
	messages := make(chan string)

	// 使用channel<-语法 发送一个新的值到通道里
	go func() {
		// for i := 0; i <= 5; i++ {
		// 	messages <- strconv.Itoa(i)
		// }
		messages <- "HugoSHEN"
	}()

	// 接收消息
	msg := <-messages
	fmt.Println(msg)
}
