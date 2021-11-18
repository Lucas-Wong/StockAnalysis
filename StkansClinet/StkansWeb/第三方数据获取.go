package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/GetOtherData", getData)

	r.Run(":9090")
}

func getData(c *gin.Context) {
	url := "http://www.baidu.com"
	response, err := http.Get(url)
	if err != nil || response.StatusCode != http.StatusOK {
		c.Status(http.StatusServiceUnavailable)
		return
	}
	body := response.Body
	contentLength := response.ContentLength
	contentType := response.Header.Get("Content-Type")
	c.DataFromReader(http.StatusOK, contentLength, contentType, body, nil)
}
