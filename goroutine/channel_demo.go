package main

import "fmt"

// 定义只读通道
func channelF1Reader(mes <-chan string) {
	msg := <-mes
	fmt.Println(msg)
}

// 定义只写通道
func channelF2Writer(mes chan<- string) {
	mes <- "blog.hugo.com"
}

func main() {
	c := make(chan string)
	fmt.Println(c)

	go channelF2Writer(c)

	msg := <-c
	fmt.Println(msg)
}
