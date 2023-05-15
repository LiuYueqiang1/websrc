package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	ID   int64
	Name string `gorm:"default:'马保郭'"`
	Age  int64
}

func main() {
	db, err := gorm.Open("mysql", "root:961024@tcp(localhost:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//定义 user 的三种方式
	//	var user User
	//	user:=User{}
	//var user = new(User)

	// 根据主键查询第一条记录
	//db.First(&user)  //把地址中的值拷贝进去
	//fmt.Println(&user)

	//// 随机获取一条记录
	//db.Take(&user)
	//fmt.Println(&user)

	// 根据主键查询最后一条记录
	//var user = new(User)
	//db.Last(user)
	//fmt.Println(*user)
	// 查询所有的记录
	//var user []User
	//db.Find(&user, "name=?", "掌门人")
	////db.Find(&user)
	//fmt.Println(user)

	//查询带有掌门人的
	//db.Where("name=?", "掌门人").First(&user)
	//fmt.Println(user)
	//var user []User

	//查询除去name为“掌门人”的
	//db.Where("name<>?", "掌门人").Find(&user)
	//fmt.Println(user)

	//db.Where("name IN (?)", []string{"掌门人", "马保郭"}).Find(&user)
	//db.Debug().Find(&user)
	//fmt.Println(user)

	//带有关键字的
	//db.Where("name LIKE ?", "%马%").Find(&user)
	//fmt.Println(user)

	//AND
	//var user []User
	//db.Where("name=? AND age >= ?", "马保郭", "70").Find(&user)
	//fmt.Println(user)

	//*******Struct&Map查询
	var user = make([]User, 10)
	//db.Where(&User{Name: "掌门人", Age: 69}).Find(&user)
	//fmt.Println(user)
	//db.Where(map[string]interface{}{"name": "掌门人", "age": 69}).Find(&user)
	//fmt.Println(user)
	// 主键的切片
	db.Where([]int64{17, 18, 19}).Find(&user)
	fmt.Println(user)
	//// SELECT * FROM users WHERE id IN (20, 21, 22);
}
