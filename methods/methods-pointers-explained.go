package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	//以指针为接收者的方法被调用时，接收者既能为值又能为指针：
	//扩大方法
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	//带指针参数的函数必须接受一个指针
	//扩大方法,第一个参数是指针类型,它的类型为*Vertex,是Vertex的指针类型
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	//主方法
	v := Vertex{3, 4}
	vPoint := &v    //取v的地址
	vPoint.Scale(2) //扩大2倍
	//v.Scale(2) //Go会将语句v.Scale(2)解释为(&v).Scale(2)。
	fmt.Println(v)
	ScaleFunc(&v, 10) //在扩大10倍
	fmt.Println(v)
}
