// *注意：* 此程序的运行环境是固定的，因此 rand.Intn 总是会返回相同的数字。
//  （要得到不同的数字，需为生成器提供不同的种子数，参见 rand.Seed。 练习场中的时间为常量，因此你需要用其它的值作为种子数。）

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favourite number is", rand.Intn(10))
}
