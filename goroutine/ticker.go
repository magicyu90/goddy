package main

import (
	"fmt"
	"time"
)

func main() {

	// 创建一个定时器tiker, 500毫秒触发一次
	ticker := time.NewTicker(time.Millisecond * 500)

	// 创建一个计时器 2秒后触发
	stopper := time.NewTimer(time.Second * 2)

	// 计数器
	var i int

	// 不断的检查
	for {
		select {
		case <-stopper.C:
			fmt.Println("结束了")
			return
		case <-ticker.C:
			i++
			fmt.Println("tick", i)
		}

	}
}
