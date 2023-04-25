package main

import "fmt"

func main() {
	var a = 10
	var b = 20
	fmt.Println(a+b)
	fmt.Println(a-b)
	fmt.Println(a*b)
	fmt.Printf("%T\n",a/b)
	fmt.Println(a%b)
	a++
	fmt.Println(a)
	fmt.Println(a == b)
	fmt.Println(a != b)
	fmt.Println(a > b)
	fmt.Println(a >= b)
	fmt.Println(a < b)
	fmt.Println(a <= b)
	if a > 30 && a < 40 {
		fmt.Println("ok1")
	}
	if a > 30 || a < 40 {
		fmt.Println("ok2")
	}
	if !(a > 30) {
		fmt.Println("ok3")
	}

	b = 30
	b += 5  // b=b+5
	fmt.Println(b)
	b *= 5  // b=b*5
	fmt.Println(b)


	var slice = []int{1,2,3,3,4}
	slice = make([]int,0)
}
