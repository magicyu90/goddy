//方法与指针重定向

package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	//定义方法
	//以指针为接收者的方法被调用时，接收者既能为值又能为指针：
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	//定义函数
	//带指针参数的函数必须接受一个指针
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2)        //使用方法扩大2倍
	ScaleFunc(&v, 10) //使用函数扩大10倍

	p := &Vertex{4, 3} //定义指针
	p.Scale(2)         //使用方法扩大2倍
	ScaleFunc(p, 10)   //使用函数扩大10倍

	fmt.Printf("v变化后是:%.2f,p变化后是:%.2f", v, *p)
}
