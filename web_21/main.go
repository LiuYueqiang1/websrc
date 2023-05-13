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

type User2 struct {
	ID   int64
	Name *string `grom:"default:'小王'"`
	Age  int64
}

func main() {
	db, err := gorm.Open("mysql", "root:961024@tcp(localhost:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.AutoMigrate(User{})  //新建表user
	db.AutoMigrate(User2{}) // 新建表 user2
	user := User{
		Name: "掌门人",
		Age:  69,
	}
	fmt.Println(db.NewRecord(user)) //查询主键是否为空  true
	db.Create(&user)                //创建user
	fmt.Println(db.NewRecord(user)) //查询主键是否为空 false
	user2 := User{
		Name: "",
		Age:  73,
	}
	db.Create(&user2)

	user3 := User2{
		Name: new(string),
		Age:  18,
	}
	db.Create(&user3)

}
