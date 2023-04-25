package main

import (
	"fmt"
	"sync"
)

var num int
var wg sync.WaitGroup
var mtx sync.Mutex

func add() {
	defer wg.Done()
	mtx.Lock()
	num += 1
	mtx.Unlock()
}

func main() {
	for i := 0; i < 1000 ; i ++ {
		wg.Add(1)
		go add()
	}
	wg.Wait()
	fmt.Println("num:", num)
}
