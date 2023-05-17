package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	ID   int64
	Name string
	Age  int64
}

func main() {
	db, err := gorm.Open("mysql", "root:961024@tcp(localhost:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//***批量删除***//
	//删除 age 为1151 的记录
	var user = make([]User, 10)
	//db.Where("age LIKE ?", 1151).Delete(user)
	//删除name 为 马保郭 的记录
	//	db.Delete(user, "name LIKE ?", "%马保郭%")

	//db.Delete(&user)
	//UPDATE users SET deleted_at="2013-10-29 10:23" WHERE id = 111;

	//删除 age 为 69 的记录
	//db.Where("age = ?", 69).Delete(&user)

	//查询
	//db.Where("age=70").Find(&user)
	//fmt.Println(user)

	// Unscoped 方法可以查询被软删除的记录 (事实上我并没有查询到被删除的记录)
	//db.Unscoped().Where("age=73").Find(&user)
	//Unscoped return all record including deleted record
	//fmt.Println(user)

	//物理删除
	db.Unscoped().Delete(&user)
}
