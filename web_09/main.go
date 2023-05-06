package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func main() {
	r := gin.Default()
	//r.LoadHTMLFiles("./websrc/web_09/templates/index.tmpl") //模板解析
	// 静态文件
	r.Static("/static", "./websrc/web_09/statics")
	//gin 框架中给模板添加自定义函数
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	r.LoadHTMLGlob("./websrc/web_09/templates/**/*")
	r.GET("posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{ //渲染模板
			"title": "liwenzhou.com",
		})
	})
	r.GET("users/posts", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts.tmpl", gin.H{ //渲染模板
			"title": "<a href='http://liwenzhou.com'>liwenzhou的博客</a>",
		})
	})
	r.Run(":9090")
}
