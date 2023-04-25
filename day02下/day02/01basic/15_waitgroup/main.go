package main

import (
	"fmt"
	"sync"
)

//第一步：定义一个计数器
var wg sync.WaitGroup

//开启两个协程
func main() {
	//100个协程的例子
	wg.Add(100)
	for i := 0; i < 100; i ++ {
		go test1()
	}
	wg.Wait()
	//wg.Add(1)
	//go test1()
	//wg.Add(1)
	//go test2()
	//wg.Wait()
	fmt.Println("主线程退出")
}

func test1() {
	//可以配合defer这样去使用
	//defer wg.Done()
	for i := 0; i < 10; i ++ {
		fmt.Println("test1() 你好golang")
	}
	wg.Done()
}

func test2() {
	for i := 0; i < 10; i ++ {
		fmt.Println("test2() 你好golang")
	}
	wg.Done()
}
