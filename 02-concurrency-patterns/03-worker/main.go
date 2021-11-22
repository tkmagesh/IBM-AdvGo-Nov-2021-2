package main

import (
	"fmt"
	"math/rand"
	"time"
	"worker-demo/worker"
)

//Any object that has a "Task()" method

var names = []string{
	"steve",
	"bob",
	"mary",
	"therese",
	"jason",
	"Magesh",
	"Ganesh",
	"Ramesh",
	"Rajesh",
	"Suresh",
}

type NamePrinter struct {
	name  string
	delay time.Duration
}

//TODO : make the tasks complete at random interval
func (np *NamePrinter) Task() {
	fmt.Println("Name Printer - Name : ", np.name)
	time.Sleep(np.delay)
}

func main() {
	p := worker.New(5)
	for idx := 0; idx < 2; idx++ {
		for _, name := range names {
			np := &NamePrinter{
				name:  name,
				delay: time.Duration(rand.Intn(len(names))) * time.Second,
			}
			p.Run(np)
		}
	}
	fmt.Println("All tasks are assigned")
	p.Shutdown()
}
