//接口与隐式实现
package main

import "fmt"

type I interface {
	M() //定义接口方法
}

type T struct {
	//定义类型
	S string
}

//此方法标识类型T实现了接口I，但我们无需显示声明此事
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	i := T{"Hugo"}
	i.M()
}
