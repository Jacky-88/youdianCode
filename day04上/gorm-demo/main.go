package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	MemberNumber int `gorm:"index"`
	CreditCards  []CreditCard `gorm:"foreignkey:UserMemberNumber;references:MemberNumber"`
}

type CreditCard struct {
	gorm.Model
	Number           string
	UserMemberNumber int
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test_db?charset=utf8&parseTime=True&loc=Local"
	//建立连接
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	//自动创建表,AutoMigrate没有建立映射关系
	//用于项目初始化自动创建表
	db.AutoMigrate(User{},CreditCard{})
	user := &User{
		MemberNumber: 155,
		CreditCards: []CreditCard{
			{Number: "0001"},
			{Number: "0002"},
		},
	}
	db.Create(user)
}
