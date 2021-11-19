package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

var ErrInterrupt = errors.New("interrupt received")
var ErrTimeout = errors.New("timeout occured")

type Runner struct {
	timeout   <-chan time.Time
	tasks     []func(int)
	complete  chan error
	interrupt chan os.Signal
}

func New(t time.Duration) *Runner {
	return &Runner{
		timeout:   time.After(t),
		complete:  make(chan error),
		interrupt: make(chan os.Signal),
	}
}

func (r *Runner) Add(task func(int)) {
	r.tasks = append(r.tasks, task)
}

func (r *Runner) Start() error {
	signal.Notify(r.interrupt, os.Interrupt)
	go func() {
		r.complete <- r.run()
	}()
	select {
	case err := <-r.complete:
		return err
	case <-r.timeout:
		return ErrTimeout
	case <-r.interrupt:
		return ErrInterrupt
	}
}

func (r *Runner) run() error {
	for id, task := range r.tasks {
		/* if r.gotInterrupt() {
			return ErrInterrupt
		} */
		task(id)
	}
	return nil
}

/* func (r *Runner) gotInterrupt() bool {
	select {
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		fmt.Println("Interrupted and exiting ")
		return true
	default:
		return false
	}
} */
