package main

//引入orm包时，需要同时引入驱动包
//"gorm.io/driver/mysql"就是gorm，mysql的驱动包
import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
//表对应的struct
type User struct {
	ID       int64   `gorm:"primary_key"`
	Username string
	Password string
}

//自定义表名
func(*User) TableName() string {
	//返回表名
	return "user_t"
}


//数据库初始化
func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test_db?charset=utf8&parseTime=True&loc=Local"
	//建立连接
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	//自动创建表,AutoMigrate没有建立映射关系
	//用于项目初始化自动创建表
	db.AutoMigrate(User{})
}
