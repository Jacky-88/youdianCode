package main

import "fmt"

func main() {
	//匿名函数func() {}()
	//recover()函数用于捕获panic异常
	//defer是在函数结束前最后执行，场景：文件关闭、db连接关闭、配合recover捕获panic。。。
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("捕获panic异常", r)
		}
	}()
	//常用用法，defer卸载建立连接或者文件打开之后
	//db.connect()
	//defer db.close()
	var mp map[string]string
	mp["username"] = "zhangsan"
}