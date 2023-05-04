package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./websrc/web_09/templates/index.tmpl") //模板解析

	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{ //渲染模板
			"title": "liwenzhou.com",
		})
	})
	r.Run(":9090")
}
