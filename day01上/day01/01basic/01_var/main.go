package _1_var

import "fmt"
var age5 = 20
func main() {
	//变量定义： var 变量名 类型 = 值
	var name string = "zhangsan"
	var age int = 21
	var isOk bool
	var age2 int
	fmt.Println(name, age, isOk, age2)
	//类型推导方式定义变量，声明时不写类型
	var age3 = 18
	//第二种方式
	age4 := 19
	fmt.Printf("%T %T\n", age3,age4)
	fmt.Println(age5)
	//一次定义多个变量
	var username, sex string
	username = "zhangsan"
	sex = "male"
	fmt.Println(username, sex)
	//
	var a, b = 1, 2
	fmt.Println(a,b)
	//常用方式，定义多个变量
	var (
		user string
		city = "beijing"
	)
	fmt.Println(user,city)
}