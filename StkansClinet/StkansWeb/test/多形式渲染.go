package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/json", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"html": "<b>Hello, Gin</b>",
		})
	})

	r.GET("/html", func(c *gin.Context) {
		c.PureJSON(http.StatusOK, gin.H{
			"html": "<b>Hello, Gin</b>",
		})
	})

	r.GET("/xml", func(c *gin.Context) {
		type Message struct {
			Name string
			Msg  string
			Age  int
		}
		info := Message{}
		info.Name = "Lucas"
		info.Msg = "Hello"
		info.Age = 30

		c.XML(http.StatusOK, info)
	})

	r.GET("/yaml", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{
			"message": "Gin 的多形式渲染",
			"status":  200,
		})
	})

	r.Run(":9090")
}
