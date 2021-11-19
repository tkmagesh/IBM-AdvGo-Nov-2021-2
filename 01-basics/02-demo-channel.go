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
	ch1 := make(chan bool)
	ch2 := make(chan bool)
	ch3 := make(chan bool)
	go f1(ch1)
	go f2(ch2)
	go f3(ch3)
	for idx := 0; idx < 3; idx++ {
		select {
		case <-ch1:
			fmt.Println("f1 done")
		case <-ch2:
			fmt.Println("f2 done")
		case <-ch3:
			fmt.Println("f3 done")
		}
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
