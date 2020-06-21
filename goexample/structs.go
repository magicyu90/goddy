//结构体

package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"hugo", 20})
	fmt.Println(person{name: "Alice", age: 30})

	fmt.Println(person{name: "Fred"})
	fmt.Println(&person{name: "Hugo", age: 29})

	s := person{name: "Hugo", age: 30}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.name)

	sp.age = 51
	fmt.Println(sp.age)
}
