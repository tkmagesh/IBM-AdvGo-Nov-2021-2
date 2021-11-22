package main

import (
	"fmt"
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
	name string
}

//TODO : make the tasks complete at random interval
func (np *NamePrinter) Task() {
	fmt.Println("Name Printer - Name : ", np.name)
	time.Sleep(2 * time.Second)
}

func main() {
	p := worker.New(5)
	for idx := 0; idx < 2; idx++ {
		for _, name := range names {
			np := &NamePrinter{name: name}
			p.Run(np)
		}
	}
	fmt.Println("All tasks are assigned")
	p.Shutdown()
}
