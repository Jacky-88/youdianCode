package main

import "fmt"

func main() {
	//切片的定义： var 变量名 []类型
	//切片没有初始化时，是nil
	var slice []string
	fmt.Println(slice)
	fmt.Println(slice == nil)
	var slice2=[]int{}
	fmt.Println(slice2 == nil)
	var slice3 = []bool{false, true}
	fmt.Println(slice3)
	fmt.Println(slice3 == nil)
	//第二种定初始化方式
	var slice4 = make([]int,0) //长度和容量
	fmt.Println(slice4 == nil)

	//引用类型，slice map和channel都是引用类型，数据传递时是引用传递
	var slice5 = []int{1,2,3}
	slice6 := slice5
	slice6[0] = 3
	slice5[1] = 5
	fmt.Println(slice5, slice6)

	//切片的长度和容量
	a := [8]int{0,1,2,3,4,5,6,7}
	fmt.Printf("%T\n", a)
	b := a[3:6]
	fmt.Printf("%T 长度%d 容量%d 值%v\n", b, len(b), cap(b),b)
	c := b[:cap(b)] //b[0:5]
	fmt.Printf("%T 长度%d 容量%d 值%v\n", c, len(c), cap(c),c)

	var slice7 = []int{1,2,3,4,5}
	//第一种遍历方式
	for i := 0; i < len(slice7); i ++ {
		fmt.Println(slice7[i])
	}
	//第二种遍历方式
	for index, value := range slice7 {
		fmt.Println(index, value)
	}

	//append添加
	var slice8 = make([]int,0)
	for i := 0; i < 10; i++ {
		slice8 = append(slice8, i)
	}
	//append添加多个
	var slice9 = []int{10,11,12}
	slice8 = append(slice8, slice9...)
	slice8 = append(slice8, 13, 14)
	fmt.Println(slice8)
	//切片中删除元素
	var slice10 = []int{30,31,32,33,34,35}
	slice10 = append(slice10[:2], slice10[3:]...)
	fmt.Println(slice10)
}