//指针接收者
//指针接收者的方法可以修改接收者指向的值（就像 Scale 在这做的）。
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale() {
	//指针接收者方法可以改变接收者本身的值，像此方法，可以直接修改v包含的值
	//没有定义返回值
	//如果使用值接收者，那么Scale方法会对原始Vertex值的副本进行操作。
	v.X = v.X * 10
	v.Y = v.Y * 10
}

func main() {
	fmt.Println("方法指针接收者")
	v := Vertex{3, 4}
	v.Scale() //v本身扩大10倍
	fmt.Println(v)
	fmt.Println(v.Abs())
}
