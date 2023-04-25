package main

import "fmt"

func main() {
	//字符串定义
	var s1 = "hello"
	//多行字符串定义
	var s2 = `
		第一行
		第二行
		第三行`
	fmt.Print(s1, s2)
}
