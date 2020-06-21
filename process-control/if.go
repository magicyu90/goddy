// 3 if 使用

package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	result := math.Sqrt(x)
	return fmt.Sprint(result)
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}
