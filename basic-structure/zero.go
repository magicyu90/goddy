//12.零值
package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

// [字符串与字节切片]

// %s 字符串或切片的无解译字节

// %q 双引号围绕的字符串，由 Go 语法安全地转义
