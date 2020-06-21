//实现 WordCount。它应当返回一个映射，其中包含字符串 s 中每个“单词”的个数。函数 wc.Test 会对此函数执行一系列测试用例，并输出成功还是失败。

package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	fmt.Println(words)
	result := make(map[string]int)
	for _, word := range words {
		elem, ok := result[word]
		if ok == false {
			result[word] = 1
		} else {
			result[word] = elem + 1
		}
	}

	return result
}

func main() {
	wc.Test(WordCount)
}
