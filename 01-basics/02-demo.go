package main

import (
	"fmt"
	"time"
)

/*
	1. using waitgroups
	2. using channels
*/
func main() {
	fmt.Println("main started")
	go f1()
	go f2()
	go f3()
	fmt.Println("main completed")
}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(10 * time.Second)
	fmt.Println("f1 completed")
}

func f2() {
	fmt.Println("f2 started")
	time.Sleep(5 * time.Second)
	fmt.Println("f2 completed")
}

func f3() {
	fmt.Println("f3 started")
	time.Sleep(7 * time.Second)
	fmt.Println("f3 completed")
}
