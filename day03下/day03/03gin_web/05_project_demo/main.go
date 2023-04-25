package main

import (
	"day03/03gin_web/05_project_demo/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routers.LoadBooks(r)
	routers.LoadUsers(r)
	r.Run(":8000")
}
