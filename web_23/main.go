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

	//var user User
	//Save()默认会更新该对象的所有字段，即使你没有赋值
	//db.First(&user)
	//user.Name = "英国大力士"
	//user.Age = 200
	//db.Save(&user)

	//	db.Model(&user).Update("name", "健身房小伙子") //将所有的都更改了

	//db.Model(&user).Where("active=?", true).Update("name", "偷袭")

	// 使用 struct 更新多个属性，只会更新其中有变化且为非零值的字段
	//db.Model(&user).Updates(User{Name: "健身房小伙子", Age: 18}) //更改了所有的

	//
	//db.Model(&user).Select("name").Updates(map[string]interface{}{"name": "嘤国大力士", "age": 18, "active": false})
	//db.Model(&user).Update(User{Name: "大力士", Age: 18})

	//更新单个属性
	//	db.Model(&user).UpdateColumn("name", "嘤国")

	//更新多个属性
	//	db.Model(&user).UpdateColumn(User{Name: "马老师", Age: 69})

	//先查询表中的第一条数据保存至user变量。
	//var user User
	//db.First(&user)
	//db.Model(&user).Update("age", gorm.Expr("age * ? + ?", 2, 100))

	//将所有的age字段做如下操作
	//var user = make([]User, 10)
	//db.Model(&user).Updates(map[string]interface{}{"age": gorm.Expr("age * ? + ?", 2, 100)})

	//将 age>10 的user做如下操作
	var user = make([]User, 10)
	db.Model(&user).UpdateColumn("age", gorm.Expr("age-?", 1)) //更新所有的age
	//db.Model(&user).Where("age>10").UpdateColumn("age", gorm.Expr("age - ?", 100))

}
