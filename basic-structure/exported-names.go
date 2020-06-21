// 在 Go 中，如果一个名字以大写字母开头，那么它就是已导出的。例如，Pizza 就是个已导出名，Pi 也同样，它导出自 math 包。
// pizza 和 pi 并未以大写字母开头，所以它们是未导出的。

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Pi)
}
