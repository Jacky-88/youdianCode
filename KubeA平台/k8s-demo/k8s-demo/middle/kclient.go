package middle

import (
	"github.com/gin-gonic/gin"
)

func KClientReverse() gin.HandlerFunc {
	return func(c *gin.Context) {
		var cluster string
		//获取请求方法
		method := c.Request.Method
		if method == "GET" {
			cluster = c.Param("cluster")
		} else {
			cluster = c.PostForm("cluster")
		}
		c.Set("cluster", cluster)
	}
}
