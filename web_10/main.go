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
	r.Static("/assets", "F:\\goland\\go_project\\go_web\\modle\\assets")
	//r.Static("/static", "F:\\goland\\go_project\\go_web\\websrc\\web_10/statics")
	//gin 框架中给模板添加自定义函数
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	//r.LoadHTMLFiles("F:\\goland\\go_project\\go_web\\modle\\index.html") //模板解析
	r.LoadHTMLGlob("F:\\goland\\go_project\\go_web\\modle\\templates/*")
	//r.GET("posts/index", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "index.tmpl", gin.H{ //渲染模板
	//		"title": "liwenzhou.com",
	//	})
	//})
	//r.GET("users/posts", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "posts.tmpl", gin.H{ //渲染模板
	//		"title": "<a href='http://liwenzhou.com'>liwenzhou的博客</a>",
	//	})
	//})
	r.GET("index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.GET("widgets.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "widgets.html", nil)
	})
	r.GET("index.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.GET("index2.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index2.html", nil)
	})
	r.Run(":9090")
}
