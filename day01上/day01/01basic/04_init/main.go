package main

import "fmt"

var a = 1

func main() {
	fmt.Println("这里是main函数")
}

func init() {
	fmt.Println("这里是init函数")
}
