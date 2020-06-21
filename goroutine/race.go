//用不同的模式，使用无缓冲的通道，在 goroutine 之间同步数据，来模拟接力比赛。在接力比赛里，4 个跑步者围绕赛道轮流跑。
//第二个、第三个和第四个跑步者要接到前一位跑步者的接力棒后才能起跑。
//比赛中最重要的部分是要传递接力棒，要求同步传递。在同步接力棒的时候，参与接力的两个跑步者必须在同一时刻准备好交接

package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	// 创建接力棒
	baton := make(chan int)
	// 计数器加1
	wg.Add(1)

	go Runner(baton)

	baton <- 1
	// 等待比赛结束
	wg.Wait()
}

//Runner 选手跑步
func Runner(baton chan int) {
	// 新选手
	var newRunner int
	// 选手等待接力棒
	runner, ok := <-baton
	if !ok {
		fmt.Printf("选手%d完成了比赛", runner)
		return
	}

	fmt.Printf("选手%d开始进行接力\n", runner)

	if runner != 4 {
		newRunner = runner + 1
		go Runner(baton)
	}
	// 模拟选手进行跑步
	time.Sleep(2 * time.Second)
	// 比赛是否结束
	if runner == 4 {
		close(baton)
		return
	}
	fmt.Printf("选手%d将接力棒交给下一位选手%d\n", runner, newRunner)
	// 放到队列
	baton <- newRunner

}
