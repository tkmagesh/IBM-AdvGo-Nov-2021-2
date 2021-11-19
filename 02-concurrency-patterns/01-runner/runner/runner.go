package runner

import (
	"errors"
	"time"
)

var ErrInterrupt = errors.New("interrupt received")
var ErrTimeout = errors.New("timeout occured")

type Runner struct {
	timeout  <-chan time.Time
	tasks    []func(int)
	complete chan error
}

func New(t time.Duration) *Runner {
	return &Runner{
		timeout:  time.After(t),
		complete: make(chan error),
	}
}

func (r *Runner) Add(task func(int)) {
	r.tasks = append(r.tasks, task)
}

func (r *Runner) Start() error {
	go func() {
		r.complete <- r.run()
	}()
	select {
	case err := <-r.complete:
		return err
	case <-r.timeout:
		return ErrTimeout
	}
}

func (r *Runner) run() error {
	for id, task := range r.tasks {
		task(id)
	}
	return nil
}
