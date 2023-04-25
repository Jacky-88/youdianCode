package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/book/:id", GetBookDetailHandler)
	r.GET("/user", GetUserDetailHandler)
	r.Run(":8000")
}

func GetBookDetailHandler(c *gin.Context) {
	bookId := c.Param("id")
	username := c.Query("name")
	c.String(http.StatusOK, fmt.Sprintf("成功获取书籍详情：%s %s", bookId, username))
}

func GetUserDetailHandler(c *gin.Context) {
	username := c.Query("name")
	c.String(http.StatusOK, fmt.Sprintf("成功获取用户详情：%s", username))
}