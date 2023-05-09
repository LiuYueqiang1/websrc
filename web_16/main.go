package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// gin重定向网页跳转
func main() {
	r := gin.Default()
	r.GET("/index", func(c *gin.Context) {
		//c.JSON(http.StatusOK, gin.H{
		//	"message": "ok",
		//})
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
	})

	r.GET("/a", func(c *gin.Context) {
		c.Request.URL.Path = "/b"
		r.HandleContext(c)
	})
	r.GET("/b", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})
	r.Run()
}
