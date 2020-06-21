package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// User 用户表
type User struct {
	Id       int    `gorm:"AUTO_INCREMENT"`
	Name     string `gorm:"size:50"`
	Age      int    `gorm:"size:3"`
	Birthday *time.Time
	Email    string `gorm:"type:varchar(50);unique_index"`
	PassWord string `gorm:"type:varchar(25)"`
}

var db *gorm.DB

func c() (i int) {
	defer func() { i++ }()
	return 1
}

// 主函数
func main() {

	// result := c()
	// fmt.Println(result)
	fmt.Printf("开始连接数据库...\n")
	db, err := gorm.Open("mysql", "#")
	if err != nil {
		//panic(err)
		fmt.Println(err)
		//return
	} else {
		fmt.Printf("连接成功\n")
	}
	defer db.Close()

	db.SingularTable(true)
	// 如果没有表
	if !db.HasTable("user") {
		db.CreateTable(&User{})
	}

	// 插入记录
	//db.Create(&User{Name: "Hugo", Age: 30, Email: "123@qq.com"})
	//db.Create(&User{Name: "Soleil", Age: 29, Email: "123@qq.com"})

	var users []User
	var user User
	// 查看插入之后数据
	fmt.Printf("插入之后数据\n")
	db.Find(&users)
	fmt.Println(users)

	// 查询第一条记录
	db.First(&user, "name=?", "Hugo")
	fmt.Println(user)

	db.Model(&user).Update("Age", 27)
	fmt.Println("更新之后的记录:", user)
}
