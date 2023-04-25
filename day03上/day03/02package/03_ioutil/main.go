package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	//读文件数据
	fileName1 := "file.txt"
	data, _ := ioutil.ReadFile(fileName1)
	fmt.Println(string(data))

	//写数据
	//fileName2 := "hello.txt"
	//s1 := "hello world"
	//ioutil.WriteFile(fileName2, []byte(s1), 0777)

	//读取reader类型的数据，常用于http response的读取
	s2 := "qweretrwy"
	r1 := strings.NewReader(s2)
	data, _ = ioutil.ReadAll(r1)
	fmt.Println(string(data))

	//读取一个目录下的所有文件和文件夹，但是只有一层
	dir, _ := os.Getwd()
	infos, _ := ioutil.ReadDir(dir)
	for _, val := range infos {
		fmt.Println(val.Name())
	}

}
