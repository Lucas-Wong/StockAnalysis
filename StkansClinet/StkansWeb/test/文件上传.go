package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	r := gin.Default()
	
	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("filename")
		if err != nil {
			c.String(http.StatusBadRequest, "上传文件错误")
		}
		dst := "D:\\Lucas\\1-Workspace\\5-Golang\\Project\\StockAnalysis\\StkansClinet\\StkansWeb\\"
		c.SaveUploadedFile(file, dst + file.Filename)
		c.String(http.StatusOK, fmt.Sprintf("%s 文件上传完成", file.Filename))
	})

	r.Run(":9090")
}
