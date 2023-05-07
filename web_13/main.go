package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 从网页中读取年月信息，或者名字，年龄等信息
func main() {
	r := gin.Default()
	r.GET("/user/:name/:age", func(c *gin.Context) {
		// 获得路径参数
		name := c.Param("name")
		age := c.Param("age")

		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})
	//http://127.0.0.1:9090/blog/2022/12
	//{"month":"12","year":"2022"}
	r.GET("blog/:year/:month", func(c *gin.Context) {
		year := c.Param("year")
		month := c.Param("month")
		c.JSON(http.StatusOK, gin.H{
			"year":  year,
			"month": month,
		})
	})
	r.Run(":9090")
}
