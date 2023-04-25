package main

import (
	"fmt"
	"os"
)

func main() {
	//获取当前目录
	fmt.Println(os.Getwd())
	// cd到某个目录
	//os.Chdir("../")
	//fmt.Println(os.Getwd())
	//创建文件夹
	//os.Mkdir("go_test", 0777)
	//删除文件夹
	//os.Remove("go_test")
	//修改文件或者文件夹的名称
	//os.Rename("go_test", "go_test123")
	os.Create("./file.txt")
	file, _ := os.OpenFile("file.txt", os.O_RDWR|os.O_APPEND, 0666)
	file.WriteString("你好 golang")

}
