package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var b = true
	fmt.Printf("%T\n", b)
	//查看占用字节
	fmt.Println(unsafe.Sizeof(b))
}
