package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("F:\\goland\\go_project\\go_web\\websrc\\web_15\\index.html")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.POST("/upload", func(c *gin.Context) {
		f, err := c.FormFile("f1")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		} else {
			//	dst := fmt.Sprintf("F:\\goland\\go_project\\go_web\\websrc\\web_15\\%s", f.Filename)
			dst := path.Join("F:\\goland\\go_project\\go_web\\websrc\\web_15", f.Filename)
			c.SaveUploadedFile(f, dst)
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "OK",
			})
		}
	})
	r.Run()
}
