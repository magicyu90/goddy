// 结构体指针
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v   // 将结构体的指针赋值p
	p.X = 1e9 //按理说需要*p.X的形式来访问，不过这么写太啰嗦了，所以语言也允许我们使用隐式间接引用，直接p.X就可以
	fmt.Println(v)
}
