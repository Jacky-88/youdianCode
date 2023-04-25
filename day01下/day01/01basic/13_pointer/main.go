package main

import "fmt"

func main() {
	//指针的3个概念
	var a = 10
	//指针地址
	fmt.Println(a,&a)
	//指针取值
	//指针类型定义
	var b *int
	b = &a
	//b的值存的是a的内存地址,&b是b本身的内存地址
	fmt.Println(b, &b)
}