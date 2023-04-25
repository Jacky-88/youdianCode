package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//求字符串长度
	var str = "this is str"
	fmt.Println(len(str))
	//字符串拼接
	var str1 = "你好"
	var str2 = "golang"
	fmt.Println(str1, str2)
	fmt.Println(str1 + ":" + str2)
	var str3 = fmt.Sprintf("%s %s", str1, str2)
	fmt.Println(str3)
	//split
	var str4 = "123,456,789"
	var arr = strings.Split(str4, ",")
	fmt.Println(arr)
	//join
	var str5 = strings.Join(arr, "-")
	fmt.Println(str5)

	//字符串遍历
	var str6 = "hello 张三"
	//第一种遍历方式,for循环出来是byte类型
	fmt.Println("------------------")
	for i := 0; i < len(str6); i++ {
		fmt.Println(str6[i])
	}
	//第二种遍历方式, for循环出来是rune类型
	fmt.Println("------------------")
	//range循环，第一个返回值是index或key，第二个是value
	for _, value := range str6 {
		fmt.Print(string(value))
	}
	//数字类型转字符串类型
	fmt.Println("\n------------------")
	//int
	var num1 int = 20
	var str7 = strconv.Itoa(num1)
	fmt.Printf("%T %s\n", str7, str7)
	//float
	//参数1：要转换的值
	//参数2：格式化类型
	//参数3：保留几个小数点
	//参数4：float的位数
	var num2 float64 = 20.113123
	var str8 = strconv.FormatFloat(num2, 'f', 2, 64)
	fmt.Printf("%T %s\n", str8, str8)

	//string转int
	var str9 = "100"
	var num3, _ = strconv.Atoi(str9)
	fmt.Printf("%T %d", num3, num3)
}
