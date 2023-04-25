package main

import "fmt"

func main() {
	//数组的定义：var 变量名 [长度]类型
	var a [5]int
	var b [3]int
	b[0] = 80
	b[1] = 100
	b[2] = 200
	fmt.Println(a, b)
	//定义时赋值
	var c = [1]string{"Tony"}
	fmt.Println(c)

	//数组在进行数据传递时，属于值传递
	var arr = [3]int{1,2,3}
	arr2 := arr
	arr2[0] = 3
	fmt.Println(arr,arr2)

	//数组的遍历
	//[...]表示不限长度的数组,但实际还是根据数组元素的数量来定义长度
	var arr3 = [...]string{"北京", "上海", "深圳"}
	//第一种遍历方式
	for i := 0; i < len(arr3); i ++ {
		fmt.Println(arr3[i])
	}
	//第二种遍历方式
	for index, value := range arr3 {
		fmt.Println(index, value)
	}
}
