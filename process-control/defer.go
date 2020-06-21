// defer 函数

package main

import "fmt"

func main() {
	defer fmt.Println("world")
	fmt.Println("hello")
}
