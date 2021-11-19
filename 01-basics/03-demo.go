package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	wg.Add(10)
	for idx := 0; idx < 10; idx++ {
		go func(i int) {
			fn(i)
		}(idx)
	}
	wg.Wait()
}

func fn(n int) {
	fmt.Println("n = ", n)
	wg.Done()
}
