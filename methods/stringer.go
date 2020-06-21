package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	//格式化字符串，本身不进行打印
	return fmt.Sprintf("姓名:%v 年龄:%v岁", p.Name, p.Age)
}

func main() {
	a := Person{"Hugo", 30}
	b := Person{"Soleil", 30}
	fmt.Println(a, b)
}
