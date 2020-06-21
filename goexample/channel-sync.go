// 通道同步
package main

import (
	"fmt"
	"time"
)

// 工作者
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true // 向channel里写入数据
}

// 主函数
func main() {

	doneChannel := make(chan bool, 1)
	go worker(doneChannel)
	fmt.Println(<-doneChannel)
}
