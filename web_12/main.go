package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 将 网页中输入框中的文件密码取出，打印到index.html网页中
func main() {
	r := gin.Default()
	r.LoadHTMLFiles("F:\\goland\\go_project\\go_web\\websrc\\web_12\\login.html", "F:\\goland\\go_project\\go_web\\websrc\\web_12\\index.html")
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	// /login post
	r.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username")         //根据html文件得到用户名
		password := c.PostForm("password")         //根据html文件得到密码
		c.HTML(http.StatusOK, "index.html", gin.H{ //加载文件
			"Name":     username,
			"Password": password,
		})
	})
	r.Run(":9090")
}
