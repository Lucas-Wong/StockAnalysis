package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("template/**/*")

	router.Static("/static", "static")

	router.GET("/", Hello)

	router.Run(":9090")
}

func Hello(ctx *gin.Context) {
	// c.String(200, "hello gin")
	ctx.HTML(200, "index/index.html", nil)
}
