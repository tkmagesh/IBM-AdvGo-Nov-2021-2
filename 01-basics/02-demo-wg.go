package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	1. using waitgroups
	2. using channels
*/
func main() {
	fmt.Println("main started")
	wg := &sync.WaitGroup{}
	wg.Add(3)
	go f1(wg)
	go f2(wg)
	go f3(wg)
	wg.Wait()
	fmt.Println("main completed")
}

func f1(wg *sync.WaitGroup) {
	fmt.Println("f1 started")
	time.Sleep(10 * time.Second)
	fmt.Println("f1 completed")
	wg.Done()
}

func f2(wg *sync.WaitGroup) {
	fmt.Println("f2 started")
	time.Sleep(5 * time.Second)
	fmt.Println("f2 completed")
	wg.Done()
}

func f3(wg *sync.WaitGroup) {
	fmt.Println("f3 started")
	time.Sleep(7 * time.Second)
	fmt.Println("f3 completed")
	wg.Done()
}
