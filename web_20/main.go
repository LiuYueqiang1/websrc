package main

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

// gorm.Model 定义 内置的
type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type User struct {
	gorm.Model   //内置的
	Name         string
	Age          sql.NullInt64 //零值
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"`        // 设置字段大小为255
	MemberNumber *string `gorm:"unique;not null"` // 设置会员号（member number）唯一并且不为空
	Num          int     `gorm:"AUTO_INCREMENT"`  // 设置 num 为自增类型
	Address      string  `gorm:"index:addr"`      // 给address字段创建名为addr的索引
	IgnoreMe     int     `gorm:"-"`               // 忽略本字段
}

// 使用方法更改表名
// 将 User 的表名设置为 `profiles`
func (User) TableName() string {
	return "profiles"
}
func main() {
	db, err := gorm.Open("mysql", "root:961024@tcp(localhost:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	//字段迁移
	db.AutoMigrate(&User{})
	//表名默认就是结构体名称的复数 users

	// 使用User结构体 创建 名为`deleted_users`的表
	db.Table("deleted_users").CreateTable(&User{})
}
