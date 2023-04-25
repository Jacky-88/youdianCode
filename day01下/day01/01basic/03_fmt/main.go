package main

import "fmt"

func main() {
	fmt.Print("zhangsan","lisi","wangwu")
	fmt.Println("zhangsan","lisi","wangwu")
	name := "zhangsan"
	age := 20
	//格式化输出
	fmt.Printf("%s,%d", name, age)
	//格式化输出，并返回字符串
	info := fmt.Sprintf("%s,%d", name, age)
	fmt.Println(info)
}
