package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

func main() {
	db, err := gorm.Open("mysql", "root:961024@tcp(localhost:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("init db failed!,err:", err)
		return
	}
	defer db.Close()

	//自动迁移
	//创建了一个名字为user_infos的表
	db.AutoMigrate(&UserInfo{}) //自动迁移为给定的模型运行自动迁移，只会添加缺失的字段，不会删除/更改当前数据

	u1 := UserInfo{1, "qimi", "男", "篮球"}
	u2 := UserInfo{2, "欧阳修", "男", "混元功法"}

	//创建记录
	db.Create(&u1)
	db.Create(&u2)
	//查询
	var u = new(UserInfo)  //将结构体的地址赋给 u
	db.First(u)            //First find first record that match given conditions, order by primary key
	fmt.Printf("%#v\n", u) //**** &main.UserInfo{ID:0x1, Name:"qimi", Gender:"男", Hobby:"篮球"}

	var uu UserInfo //将结构体的数值赋给 uu
	db.Find(&uu, "hobby=?", "混元功法")
	fmt.Printf("%#v\n", uu) //**** main.UserInfo{ID:0x2, Name:"欧阳修", Gender:"男", Hobby:"混元功法"}

	//更新
	db.Model(&u).Update("hobby", "唱跳rap")
	//删除
	db.Delete(&u) //故此处删掉的u为 ID：01的
}
