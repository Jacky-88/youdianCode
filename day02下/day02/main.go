package main

import (
	encapsulation "day02/01basic/09_encapsulation"
	"fmt"
)

func main() {
	//s1 := &encapsulation.Student{
	//	Name: "李四",
	//	score: 50,
	//}
	s := encapsulation.NewStudent("张三", 80)
	fmt.Println(s)
	s.SetScore(90)
	fmt.Println(s.GetScore())
}
