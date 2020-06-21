//底层值为 nil 的接口值
package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I
	var t *T
	i = t
	describe(i)
	t.M()

	i = &T{"Hugo"}
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v,%T)\n", i, i)
}
