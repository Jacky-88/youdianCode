package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
func main() {
	//Default返回一个默认的路由引擎
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		//输出json结果给调用方
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"code":http.StatusOK,
			"msg":"gin...",
		})
	})
	r.GET("/xml", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "<p>该项目是gin框架的学习笔记</p>",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}