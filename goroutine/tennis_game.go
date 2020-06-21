// 这个示例程序展示如何用无缓冲的通道来模拟
// 2 个goroutine 间的网球比赛

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// wg 用来等待程序结束
var wg sync.WaitGroup

func main() {

	// 生成不同随机种子数
	rand.Seed(time.Now().Unix())

	court := make(chan int)
	// 计数器加2表示有两个goroutine
	wg.Add(2)

	go player("hugo", court)
	go player("soleil", court)

	court <- 1

	// 等待游戏结束
	wg.Wait()

}

func player(name string, court chan int) {

	// 函数退出时告诉main函数工作已经完成
	defer wg.Done()

	for {
		ball, ok := <-court
		if !ok {
			// 如果通道被关闭，证明选手赢了
			fmt.Printf("Player %s Won\n", name)
			return
		}

		randomNum := rand.Intn(100)
		if randomNum%11 == 0 {
			fmt.Printf("Player %s missed the game at %d times with num:%d \n", name, ball, randomNum)
			//关闭通道
			close(court)
			return
		}

		// 显示击球数信息
		fmt.Printf("Player %s hit ball %d times with num:%d\n", name, ball, randomNum)
		// 击球数加1
		ball++
		// 将球击向对方
		court <- ball
	}
}
