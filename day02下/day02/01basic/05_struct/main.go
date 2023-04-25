package main

import "fmt"

//结构体定义  type关键字  结构体名字 struct关键字 {属性..}
type Person struct {
	Name string
	City string
	Age int
}

func main() {
	//结构体实例化
	//方法1 值类型
	var p1 Person
	p1.Name = "张三"
	p1.City = "北京"
	p1.Age = 18
	fmt.Println(p1)
	//方法2 指针类型（推荐）
	var p2 = new(Person)
	p2.Name = "李四"
	p2.City = "上海"
	p2.Age = 20
	fmt.Println(p2)
	//方法3 指针类型（推荐）
	//var p3 = &Person{}
	var p3 = &Person{
		Name: "王五",
		City: "深圳",
		Age:  30,
	}
	fmt.Println(p3)
}
