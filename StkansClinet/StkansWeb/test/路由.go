package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main1() {
	r := gin.Default()
	//r.GET("/get", getMsg)
	//r.POST("/post", postMsg)
	r.GET("/test", redirect)
	r.GET("/TestRedirect", testRedirect)
	r.GET("/redirect3", func(c *gin.Context) {
		c.Request.URL.Path = "/TestRedirect"
		r.HandleContext(c)
	})
	r.Run(":9090")
}

func getMsg(c *gin.Context) {
	name := c.Query("name")
	// c.String(http.StatusOK, "欢迎您：%s", name)
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,  //状态
		"msg":  "返回信息",         //描述信息
		"data": "欢迎您: " + name, //数据
	})
}

func postMsg(c *gin.Context) {
	name := c.DefaultPostForm("name", "Gin")
	formName, b := c.GetPostForm("name")
	fmt.Println(formName, b)
	c.JSON(http.StatusOK, "欢迎您："+name)
}

func redirect(c *gin.Context) {
	url := "http://www.baidu.com"
	c.Redirect(http.StatusMovedPermanently, url)
}

func testRedirect(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "成功响应",
	})
}
