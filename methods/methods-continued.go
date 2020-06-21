// 可以为非结构体类型声明方法

package main

import (
	"fmt"
	"math"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {

	fmt.Println("methods-continued...")
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
