package main

import (
	"bookManage/dao"
	"bookManage/router"
)

func main() {
	dao.InitMysql()
	r := router.InitRouter()
	r.Run(":8000")
}
