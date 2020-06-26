package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

// 主函数
func main() {

	var wg sync.WaitGroup

	var urls = []string{
		"http://www.github.com/",
		"https://www.qiniu.com/",
		"https://www.sina.com/",
	}
	// for循环遍历
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {

			// 执行完通知
			defer wg.Done()

			// 使用http访问提供的地址
			resp, err := http.Get(url)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			fmt.Printf("网站:%s,内容长度:%d\n", url, len(body))
		}(url)
	}
	wg.Wait()

	fmt.Printf("game over!")

}
