// switch 是编写一连串 if - else 语句的简便方法。它运行第一个值等于条件表达式的 case 语句。
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs go")
	var systemName = runtime.GOOS
	switch os := systemName; os {
	case "darwin":
		fmt.Println("os x.")
	case "linux":
		fmt.Println("linux")
	default:
		fmt.Println("\n %s.", os)
	}
}
