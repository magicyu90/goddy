package main

import "fmt"

// ping通道 发送数据
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// pings通道 接收数据 pongs通道 发送数据
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings //从通道里接收
	pongs <- msg   //发送到通道里
}

// 主函数
func main() {
	pings := make(chan string, 2) //ping通道
	pongs := make(chan string, 2) //pong通道

	ping(pings, "passed message1")
	ping(pings, "passed message2")
	pong(pings, pongs)
	pong(pings, pongs)
	fmt.Println(<-pongs)
	fmt.Println(<-pongs)
}
