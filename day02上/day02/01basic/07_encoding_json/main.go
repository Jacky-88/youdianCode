package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	ID int
	Gender string
	name string //私有属性不能被json包访问
	Sno string
}

func main() {
	var s1 = &Student{
		ID:     1,
		Gender: "男",
		name:   "张三",
		Sno:    "s0001",
	}
	fmt.Println(s1)
	//序列化
	strByte, err := json.Marshal(s1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(strByte))
	//反序列化
	str := `{"ID":3,"Gender":"女","Sno":"s0003"}`
	var s2 = &Student{}
	fmt.Println(s2)
	err = json.Unmarshal([]byte(str), s2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s2)
}
