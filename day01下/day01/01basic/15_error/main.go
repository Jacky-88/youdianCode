package main

import (
	"fmt"
	"net/http"
)

func main() {
	//第一种方式判断error
	resp, err := http.Get("http://example13452345.com")
	if err != nil {
		//日志输出
		//panic恐慌
		//return上一层 API开发常用
		//fmt打印
		//...
		fmt.Println(err)
		return
	}
	fmt.Println(resp)
	//第二种方式
	if resp1, err := http.Get("http://example13452345.com"); err != nil {
		fmt.Println(err)
		fmt.Println(resp1)
		return
	}

}
