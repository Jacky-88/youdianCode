package main

import "fmt"

//父结构体
type Person struct {
	Name string
	Age  int
	Sex  string
}

func(p *Person) SayHello() {
	fmt.Println(p)
	fmt.Println("this is from Person")
}
//子结构体
type Student struct {
	School string
	Person
}
//一般情况下不使用golang继承，用嵌套结构体替代，因为嵌套结构体比较简单，易读
//type Student struct {
//	School string
//	People *Person
//}

func main() {
	stu := &Student{
		School: "middle",
	}
	//模拟继承的关系，帮我们自动处理
	//stu.Name = stu.Person.Name
	//stu.SayHello() = stu.Person.SayHello()
	//属性和方法都能继承
	stu.Name = "Leo"
	stu.Age = 30
	stu.Sex = "Male"
	fmt.Println(stu)
	stu.SayHello()
}
