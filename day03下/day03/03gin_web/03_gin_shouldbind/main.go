package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Login struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func LoginHandler(c *gin.Context) {
	login := &Login{}
	//GET请求一般使用上面的form标签
	//application/json这种请求，POST PUT DELETE都是用json标签
	if err := c.ShouldBind(&login); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "绑定参数失败" + err.Error(),
			"data": nil,
			"code": 90400,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "接口请求成功",
		"data": login,
		"code": 90200,
	})
}

func main() {
	r := gin.Default()
	r.POST("/login", LoginHandler)
	r.Run(":8000")
}