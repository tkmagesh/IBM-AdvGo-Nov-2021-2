package pool

import (
	"errors"
	"fmt"
	"io"
	"sync"
)

var ErrInvalidPoolSize = errors.New("invalid pool size")
var ErrPoolClosed = errors.New("pool has been closed")

type Pool struct {
	factory   func() (io.Closer, error)
	resources chan io.Closer
	closed    bool
	mutex     sync.Mutex
}

func New(factory func() (io.Closer, error), size int) (*Pool, error) {
	if size <= 0 {
		return nil, ErrInvalidPoolSize
	}

	return &Pool{
		factory:   factory,
		resources: make(chan io.Closer, size),
	}, nil
}

func (p Pool) Acquire() (io.Closer, error) {
	select {
	case r, ok := <-p.resources:
		if !ok {
			return nil, ErrPoolClosed
		}
		fmt.Println("Acquire : From pool")
		return r, nil
	default:
		fmt.Println("Acquire : From factory")
		return p.factory()
	}
}

func (p *Pool) Release(r io.Closer) error {

	p.mutex.Lock()
	defer p.mutex.Unlock()

	if p.closed {
		r.Close()
		return ErrPoolClosed
	}

	select {
	case p.resources <- r:
		fmt.Println("Release : In pool")
		return nil
	default:
		fmt.Println("Release : Discard the resource")
		return r.Close()
	}
}

func (p *Pool) Close() {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if p.closed {
		return
	}

	p.closed = true

	close(p.resources)
	for r := range p.resources {
		r.Close()
	}
}
