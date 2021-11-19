package main

import (
	"fmt"
	"time"
)

func main() {
	//using a non-buffered channel
	ch := make(chan int)

	//using a buffered channel
	//ch := make(chan int, 1)
	fmt.Println("main started")
	fmt.Println("initiating the fn goroutine")
	go fn(ch)
	fmt.Println("start - sleep for 20 seconds")
	time.Sleep(20 * time.Second)
	fmt.Println("end - sleep for 20 seconds")
	fmt.Println("Attempting to read data from the channel")
	no := <-ch
	fmt.Println("no =", no)
	fmt.Println("main finished")
}

func fn(ch chan int) {
	fmt.Println("fn started")
	fmt.Println("Attempting to write data to the channel")
	ch <- 100
	fmt.Println("successfully written data to the channel")
	fmt.Println("fn finished")
}
