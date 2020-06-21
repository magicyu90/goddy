// 通道缓冲
package main

import "fmt"

// 主函数
func main() {

	// 创建通道，最多缓冲2个
	messages := make(chan string, 2)
	// 发送消息
	messages <- "buffered"
	// 发送消息
	messages <- "channel"

	fmt.Println(<-messages) // 接收消息
	fmt.Println(<-messages) // 接收消息
	//fmt.Println(<-messages) // 接收消息
}
