package main

import (
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

var DB *gorm.DB

// 初始化数据库
func initMySQL() (err error) {
	dsn := "root:961024@tcp(localhost:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	return DB.DB().Ping()
}

func main() {
	//创建数据库
	err := initMySQL()
	if err != nil {
		panic(err)
	}
	defer DB.Close()
	//模型绑定
	DB.AutoMigrate(&ToDo{})
	//加载中间件
	r := gin.Default()
	//加载静态文件
	r.Static("/static", "F:\\goland\\go_project\\go_web\\websrc\\web_25\\static")

	//模板解析
	r.LoadHTMLGlob("F:\\goland\\go_project\\go_web\\websrc\\web_25\\tmlplates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// v1
	v1Group := r.Group("v1")
	//Group命令用于创建新的路由器组。您应该添加所有具有相同中间件或相同路径前缀的路由。
	//例如，可以对使用公共中间件进行授权的所有路由进行分组。

	//待办事项
	//添加
	{
		v1Group.POST("/todo", func(c *gin.Context) {
			//前端页面填写待办事项 点击提交 发送请求到此处
			var todo ToDo
			// 将 页面上的值放到todo中
			c.BindJSON(&todo)
			// 1、 从请求中把数据拿出来
			// 2、存入数据
			//err = DB.Create(&todo).Error
			//if err!= nil {
			//}
			// 3、返回响应
			if err = DB.Create(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"code": 2000,
					"msg":  "success",
					"data": todo,
				})
			}
		})
		// 查看所有的待办事项
		v1Group.GET("/todo", func(c *gin.Context) {
			//查询todo这个表里的所有数据
			var todoList []ToDo
			if err = DB.Find(&todoList).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
			} else {
				c.JSON(http.StatusOK, todoList)
			}
		})
		// 修改某一个待办事项
		v1Group.PUT("/todo/:id", func(c *gin.Context) {
			// 拿到要修改的id
			id, ok := c.Params.Get("id")
			//Get返回与给定名称匹配的第一个Param的值和一个布尔值true。
			//如果没有找到匹配的参数，则返回一个空字符串和一个布尔值false。
			if !ok {
				c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
				return
			}
			// 根据id 匹配数据
			var todo ToDo
			if err = DB.Where("id=?", id).First(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
				return
			}
			// 修改
			c.BindJSON(&todo)
			if err = DB.Save(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, todo)
			}
		})
		v1Group.DELETE("/todo/:id", func(c *gin.Context) {
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
				return
			}
			if err = DB.Where("id=?", id).Delete(ToDo{}).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, gin.H{id: "deleted"})
			}
		})
	}

	r.Run()
}
