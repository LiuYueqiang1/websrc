package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/json", func(c *gin.Context) {
		// 方法1：使用map
		//data := map[string]interface{}{
		//	"name":    "小王子",
		//	"age":     18,
		//	"message": "hello world!",
		//}
		data := gin.H{"name": "小王子", "age": 18, "message": "hello world!"}
		//H is a shortcut for map[string]interface{}
		c.JSON(http.StatusOK, data)
	})

	// 方法2：结构体，灵活使用tag对结构体字段做定制化操作
	type msg struct {
		Name    string `json:"name"`
		Message string `json:"message"`
		Age     int    `json:"age"`
	}
	r.GET("/newjson", func(c *gin.Context) {
		data := msg{
			"马老师",
			"我打一个连五鞭,发生甚摸事了",
			69,
		}
		c.JSON(http.StatusOK, data)
	})
	r.Run()
}
