// 6 if 和 else
package main

import (
	"fmt"
	"math"
)

// 平方方法
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g>=%g\n", v, lim)
	}
	return lim
}

func main() {
	fmt.Println(pow(3, 2, 10), pow(3, 3, 20))
}
