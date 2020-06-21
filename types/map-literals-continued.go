// 映射的文法（续）

package main

import "fmt"

type VertexNew struct {
	Lat, Long float64
}

var m = map[string]VertexNew{
	"Bell Labs": {28.123, 45.899},
	"Google":    {37.032, -123.273},
}

func main() {
	fmt.Println(m)
}
