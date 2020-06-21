package main

import "fmt"

func main() {
	var i interface{} = "Hello hugo"
	s := i.(string)
	fmt.Println(s)
	fmt.Printf("%v,%T\n", s, s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)
}
