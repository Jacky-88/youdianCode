package main

import "fmt"

func main() {
	//定义一个管道，10个int数据
	intChan := make(chan int, 10)
	for i := 0; i < 10; i ++ {
		intChan <- i
	}
	//定义一个管道，5个string数据
	strChan := make(chan string, 5)
	for i := 0; i < 5; i ++ {
		strChan <- fmt.Sprintf("hello %d", i)
	}
	//使用select来获取channel里面的数据的时候，不需要关闭channel
	//当监听的channel有数据时，则走case，没有数据时，则走default
	//一般select监听channel时会配合for循环一起使用
	for {
		select {
		case v := <- intChan:
			fmt.Printf("从intChan读取的数据%d\n", v)
		case v := <- strChan:
			fmt.Printf("从strChan读取的数据%s\n", v)
		default:
			fmt.Println("数据读取完毕")
			return
		}
	}
}
