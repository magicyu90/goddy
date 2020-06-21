package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	res := n * fact(n-1)
	fmt.Printf("n:%v,当前res:%v\n", n, res)
	return res
}

func main() {
	fmt.Println(fact(7))
}
