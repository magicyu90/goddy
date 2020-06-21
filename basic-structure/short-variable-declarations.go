//短变量声明 := 结构不能在函数外使用。

package main

import "fmt"

func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "No"
	fmt.Println(i, j, k, c, python, java)
}
