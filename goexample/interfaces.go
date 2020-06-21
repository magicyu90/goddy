// 接口
package main

import (
	"fmt"
	"math"
)

// 几何学接口
type geometry interface {
	area() float64
	perim() float64
}

// 矩形类
type rect struct {
	width, height float64
}

// 圆类
type circle struct {
	radius float64
}

// 面积方法接口实现
func (r rect) area() float64 {
	return r.width * r.height
}

// 周长方法接口实现
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// 圆面积方法接口实现
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// 圆周长方法接口实现
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// 测量
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

// 主函数
func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}
	measure(r)
	measure(c)
}
