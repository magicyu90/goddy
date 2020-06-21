//通道选择器
package main

import (
	"fmt"
	"time"
)

func main() {

	//创建通道
	c1 := make(chan string) //通道1
	c2 := make(chan string) //通道2

	// go协程 处理通道
	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()

	// go协程 处理通道
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
