package main

import "fmt"

type rect struct {
	width, height int
}

// 面积 方法指针
func (r *rect) area() int {
	return r.width * r.height
}

// 周长 方法函数
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}
	fmt.Println("area", r.area())
	fmt.Println("perim", r.perim())

	rp := &r
	fmt.Printf("rp类型:%T", rp)
	fmt.Println("area", rp.area())
}
