package main

import "fmt"

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {

	a := 1                               //1
	b := 2                               //2
	defer calc("1", 1, calc("10", a, b)) //3
	a = 0                                //4
	defer calc("2", a, calc("20", a, b)) //5
	b = 1                                //6

	for i := 1; i < 10; i++ {
		fmt.Printf("i的地址:%p,i的值:%v\n", &i, i)
	}
}

// 1.先执行第三行的calc("10",a,b)=> calc("10",1,2)=3
// 2.执行defer calc("1",1,3) 存放到延迟队列中
// 3.执行calc("20", a, b)=> calc("20",0,2)=>2
// 4.执行defer calc("2",0,2) 存放到延迟队列中
// 5.后进先出 先执行defer calc("2",0,2) 再执行defer calc("1",1,3)
