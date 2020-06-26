package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int64
	mutex   sync.Mutex
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go incCounter(2)
	go incCounter(3)
	wg.Wait()
	fmt.Println(counter)
}
func incCounter(id int) {
	defer wg.Done()
	for count := 0; count < id; count++ {
		//同一时刻只允许一个goroutine进入这个临界区
		mutex.Lock()
		{
			value := counter
			runtime.Gosched()
			value++
			counter = value
		}
		mutex.Unlock() //释放锁，允许其他正在等待的goroutine进入临界区
	}
}
