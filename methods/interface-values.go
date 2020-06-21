//接口值
package main

import (
	"fmt"
	"math"
)

//定义接口I
type I interface {
	M()
}

//定义结构体T
type T struct {
	S string
}

func (t T) M() {
	fmt.Println(t.S)
}

//定义floatF
type F float64

func (f F) M() {
	fmt.Println(f)
}

func describe(i I) {
	fmt.Printf("(%v,%T)\n", i, i)
}

func main() {
	var i I
	i = T{"hugo"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()

}
