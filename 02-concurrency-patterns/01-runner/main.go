package main

import (
	"fmt"
	"os"
	"runner-demo/runner"
	"time"
)

func main() {
	/*
		initialize runner with a timeout
		Assign multiple tasks to the runner
		Start the runner
		if all the tasks are completed with the given time, report "success"
		if the tasks are not completed with the given time, report "timeout"
		if the execution is interruped by os interrupt
	*/
	fmt.Printf("Process # %d started \n", os.Getpid())
	timeout := 15 * time.Second
	r := runner.New(timeout)
	r.Add(createTask(3))
	r.Add(createTask(5))
	r.Add(createTask(8))

	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrTimeout:
			fmt.Println("tasks Timedout")
		case runner.ErrInterrupt:
			fmt.Println("interrupt received")
		}
	}

	fmt.Println("Processor ended")
}

func createTask(t int) func(int) {
	return func(id int) {
		fmt.Printf("Processor - Task #%d\n", id)
		time.Sleep(time.Duration(t) * time.Second)
	}
}
