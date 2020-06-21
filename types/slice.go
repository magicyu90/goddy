// 切片
package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 4, 6, 7, 9}
	var s []int = primes[1:4]
	fmt.Println(s)
}
