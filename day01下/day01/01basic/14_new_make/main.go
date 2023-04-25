package main

import "fmt"

func main() {
	//make，初始化申请一片内存空间
	//常用场景：slice map chan
	var a = make([]int,0)
	a = append(a, 1)
	fmt.Printf("%T %v\n", a, a)
	//new，初始化申请一片内存空间，并返回指针类型
	//struct结构体
	var b = new([]int)
	*b = append(*b, 1)
	fmt.Printf("%T %v", b, b)
}
