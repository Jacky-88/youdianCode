package main

import "fmt"

type Person struct {
	Name string
	Age int
}
//接收器，能用指针类型就用指针类型

//值类型接收器
//接收器接收的是结构体的实例化对象
//p是接收器变量名，Person是接收器类型
func(p Person) printInfo() {
	fmt.Printf("姓名:%s,年龄:%d \n", p.Name, p.Age)
}

//指针类型接收器
//改变接收器的源值
func(p *Person) setInfo(name string, age int) {
	p.Name = name
	p.Age = age
	fmt.Println("setInfo内", p)
}
//值类型接收器，copy一个接收器，不会改变源值
//func(p Person) setInfo(name string, age int) {
//	p.Name = name
//	p.Age = age
//	fmt.Println("setInfo内", p)
//}

func main() {
	//结构体实例化对象
	p1 := &Person{
		Name: "张三",
		Age:  18,
	}
	p1.printInfo()
	fmt.Println(p1)
	p1.setInfo("李四", 20)
	fmt.Println(p1)
}
