package main

import "fmt"

func main() {
	//第一种,变量在if条件判断之外,变量作用域在整个main函数中
	score := 65
	if score >= 90 {
		fmt.Println("A")
	} else if (score > 75) {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}

	//第二种，变量在if条件判断中，变量作用域在if函数中
	if score := 65; score >= 90 {
		fmt.Println("A")
	} else if (score > 75) {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}

}
