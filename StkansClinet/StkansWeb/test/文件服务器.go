package main

import "github.com/gin-gonic/gin"

func main(){
	r := gin.Default()

	r.GET("/file", fileServer)

	r.Run(":9090")
}

func fileServer(c *gin.Context){
	path := "D:\\Lucas\\1-Workspace\\5-Golang\\Project\\StockAnalysis\\StkansClinet\\StkansWeb\\"
	filename := path + c.Query("name")
	c.File(filename)
}
