package main

import "fmt"

//空接口作为函数的参数  interface{}
func show(a ...interface{})  {
	fmt.Printf("值：%v 类型：%T \n", a, a)
}

func main() {
	show(20, "ssss")
	show("你好golang")
	slice := []int{1,2,3,4}
	show(slice)
	//空接口实现切片
	var slice2 = []interface{}{"张三", 20, true, 32.2}
	fmt.Println(slice2)
	//空接口实现map
	var stuInfo = make(map[string]interface{})
	stuInfo["name"] = "张三"
	stuInfo["age"] = 19
	stuInfo["married"] = false
	fmt.Println(stuInfo)
}
