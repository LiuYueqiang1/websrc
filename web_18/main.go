package main

// gin中间件
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func indexHandler(c *gin.Context) {
	fmt.Println("index")
	c.JSON(http.StatusOK, gin.H{
		"msg": "index",
	})
}

// 定义一个gin中间件 ，统计请求处理函数的耗时
func m1(c *gin.Context) {
	fmt.Println("m1 in ...")
	//计时
	start := time.Now()
	c.Next()                  //调用后续的处理函数
	cost := time.Since(start) //统计所有处理函数花费的时间
	fmt.Println("cost:", cost)
	fmt.Println("m1 out ...")
}
func m2(c *gin.Context) {
	fmt.Println("m2 in ...")
	//c.Next()	//调用后续的处理函数
	c.Abort() //阻止调用后续函数
	fmt.Println("m2 out ...")
	//	m1 in ...
	//	m2 in ...
	//	m2 out ...
	//cost: 711.6µs
	//	m1 out ...

	//return
	//m1 in ...
	//m2 in ...
	//cost: 1.2252ms
	//m1 out ...
}
func main() {
	r := gin.Default()
	//r.GET("/index", m1, indexHandler)
	r.Use(m1, m2) //全局注册中间件m1,m2
	r.GET("/index", indexHandler)
	r.GET("/shop", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "shop",
		})
	})
	r.GET("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "user",
		})
	})
	r.Run(":9090")
}
