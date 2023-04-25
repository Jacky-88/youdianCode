package main

import "fmt"

func main() {
	//map的定义: var 变量名=make(map[key类型]value类型)
	//第一种方式，常用
	var mp = make(map[string]string)
	fmt.Println(mp)
	//第二种方式
	var mp2 = map[string]string{}
	mp2["url"] = "www.baidu.com"
	fmt.Println(mp2)
	//第三种方式
	var mp3 map[string]string
	//mp3["url"] = "www.baidu.com"
	fmt.Println(mp3)

	//判断某个key是否存在
	var userInfo = map[string]string{
		"username": "zhangsan",
		"password": "123456",
	}
	fmt.Println(userInfo)
	//第一个返回值 Value
	//第二个返回值 key是否存在 bool
	v, ok := userInfo["username"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("key不存在")
	}

	//delete()，内置函数，用于从map中删除一组键值对
	//第一个参数 map
	//第二个参数 key
	delete(userInfo, "password")
	fmt.Println(userInfo)

	//map的遍历
	var scoreMap = map[string]int {
		"zhangsan": 60,
		"lisi": 70,
		"wangwu": 80,
	}
	//key value都遍历
	for key,value := range scoreMap {
		fmt.Println(key, value)
	}
	//只遍历value
	for _,value := range scoreMap {
		fmt.Println(value)
	}
	//只遍历key
	for key := range scoreMap {
		fmt.Println(key)
	}
}
