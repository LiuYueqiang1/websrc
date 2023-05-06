package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/web", func(c *gin.Context) {
		name := c.Query("query")
		age := c.Query("age")
		//name := c.DefaultQuery("query", "somebody") //取不到就用指定的默认值

		//name, ok := c.GetQuery("query")
		//if !ok {
		//	//取不到
		//	name = "somebody"
		//}

		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})
	r.Run()
}

//http://127.0.0.1:8080/web?query=筱往&age=18
