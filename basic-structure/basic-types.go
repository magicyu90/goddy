package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

// 主函数
func main() {
	fmt.Printf("Type:%T value:%v\n", ToBe, ToBe)
	fmt.Printf("Type:%T value:%v\n", MaxInt, MaxInt)
	fmt.Printf("Type:%T value:%v\n", z, z)
}

// %v 相应值的默认格式。在打印结构体时，“加号”标记（%+v）会添加字段名

// %#v 相应值的 Go 语法表示

// %T 相应值的类型的 Go 语法表示

// %% 字面上的百分号，并非值的占位符
