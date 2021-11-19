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
	done := make(chan bool)
	go f1(done)
	go f2(done)
	go f3(done)
	for idx := 0; idx < 3; idx++ {
		<-done
	}

	fmt.Println("main completed")
}

func f1(done chan bool) {
	fmt.Println("f1 started")
	time.Sleep(10 * time.Second)
	fmt.Println("f1 completed")
	done <- true

}

func f2(done chan bool) {
	fmt.Println("f2 started")
	time.Sleep(5 * time.Second)
	fmt.Println("f2 completed")
	done <- true

}

func f3(done chan bool) {
	fmt.Println("f3 started")
	time.Sleep(7 * time.Second)
	fmt.Println("f3 completed")
	done <- true

}
