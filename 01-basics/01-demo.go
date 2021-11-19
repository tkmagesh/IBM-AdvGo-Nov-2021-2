package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//using a non-buffered channel
	ch := make(chan int)
	//done := make(chan bool)
	wg := &sync.WaitGroup{}
	//using a buffered channel
	//ch := make(chan int, 1)
	fmt.Println("main started")
	fmt.Println("initiating the fn goroutine")
	wg.Add(1)
	go fn(ch, wg)
	fmt.Println("start - sleep for 20 seconds")
	time.Sleep(5 * time.Second)
	fmt.Println("end - sleep for 20 seconds")
	fmt.Println("Attempting to read data from the channel")
	no := <-ch
	fmt.Println("no =", no)
	wg.Wait()
	fmt.Println("main finished")

}

func fn(ch chan int, wg *sync.WaitGroup) {
	fmt.Println("fn started")
	fmt.Println("Attempting to write data to the channel")
	ch <- 100
	fmt.Println("successfully written data to the channel")
	fmt.Println("fn finished")
	wg.Done()
}
