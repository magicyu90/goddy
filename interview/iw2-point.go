package main

import "fmt"

// 学生类
type student struct{
	Name string
	Age int
}

// 传递学生类
func pase_student() map[string]*student{

	m:= make(map[string]*student)

	stus := []student{
		{Name: "Hugo",Age:30},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	// for遍历时，变量stu指针不变，每次遍历仅进行struct值拷贝
	for _,stu:= range stus{
		fmt.Println(&stu)
		m[stu.Name]= &stu
	}

	// 可以取数组中原始值的指针
	for i, _ := range stus {
		stu:=stus[i]
		m[stu.Name] = &stu
	}
	return m
}

// 主函数
func main(){
	fmt.Printf("hello\n")
	students := pase_student()
	fmt.Println(students)
	for k,v:= range students{
		fmt.Printf("key=%s,value=%v \n",k,v)
	} 
}

