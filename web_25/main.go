package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

type ToDo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

var db *gorm.DB

func initMysql() (err error) {
	dsn := "root:961024@tcp(localhost:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	return
}

func main() {
	//连接数据库
	err := initMysql()
	if err != nil {
		fmt.Println("init mysql failed!,err:", err)
		return
	}

	//模型绑定
	//新建表todos
	db.AutoMigrate(&ToDo{})

	r := gin.Default()
	//加载静态文件
	r.Static("/static", "F:\\goland\\go_project\\go_web\\websrc\\web_25\\static")
	//加载文件
	r.LoadHTMLGlob("F:\\goland\\go_project\\go_web\\websrc\\web_25\\tmlplates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	v1Group := r.Group("v1")
	{
		v1Group.POST("/todo", func(c *gin.Context) {
			var todo ToDo
			// 1、 从请求中把数据拿出来
			c.BindJSON(&todo)

			//2、存入数据
			err := db.Create(&todo).Error
			if err != nil {
				fmt.Println("create err")
				return
			} else {
				//3、返回响应
				c.JSON(http.StatusOK, gin.H{
					"msg":    "2000",
					"status": "ok",
				})
			}
		})

		v1Group.GET("/todo", func(c *gin.Context) {
			var todo []ToDo
			db.Find(&todo)
			//显示数据
			c.JSON(http.StatusOK, todo)
		})
		v1Group.PUT("/todo/:id", func(c *gin.Context) {
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{
					"errors": "无效的id",
				})
				return
			}
			var todo ToDo
			//******************************************************//
			// 查询要修改的id，这个非常重要
			//如果没有这一项的话无法匹配到要修改的项，则会在每次更新的时候save两条空数据
			//First(&todo)而不是First(&ToDo)，不然会更改整个结构体，同样增加两条新数据
			err := db.Where("id = ?", id).First(&todo).Error
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"errors": "无效的id",
				})
				return
			} else {
				c.JSON(http.StatusOK, gin.H{
					"msg":    "2000",
					"status": "ok",
				})
			}

			// 从请求中将数据拿出来存入结构体todo中
			c.BindJSON(&todo)
			// 更新数据
			err = db.Save(&todo).Error
			if err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, todo)
			}
		})
		v1Group.DELETE("/todo/:id", func(c *gin.Context) {
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{
					"errors": "无效的id",
				})
				return
			}
			err := db.Where("id = ? ", id).Delete(ToDo{}).Error
			if err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, gin.H{id: "deleted"})
			}
		})
	}
	r.Run()
}
