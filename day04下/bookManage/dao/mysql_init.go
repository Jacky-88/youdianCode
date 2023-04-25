package dao

import (
	"bookManage/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
//声明全局变量
var DB *gorm.DB

func InitMysql() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/books?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("初始化数据库失败" + err.Error())
	}
	DB = db
	fmt.Println("初始化数据库成功")
	//自动创建表
	if err := DB.AutoMigrate(model.User{}, model.Book{}); err != nil {
		fmt.Println("自动创建表失败:", err)
	}
}