// 5 if 的简短语句

package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	v := math.Pow(x, n)
	if v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(pow(3, 2, 10), pow(3, 3, 20))
}
