package main

import (
	"fmt"
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

	timeout := 20 * time.Second
	r := runner.New(timeout)
	r.Add(createTask())
	r.Add(createTask())
	r.Add(createTask())

	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrTimeout:
			// timeout
		case runner.ErrInterrupt:
			// interrupt
		}
	}

	fmt.Println("Processor ended")
}

func createTask() func(int) {
	return func(id int) {
		fmt.Printf("Processor - Task #%d\n", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}
