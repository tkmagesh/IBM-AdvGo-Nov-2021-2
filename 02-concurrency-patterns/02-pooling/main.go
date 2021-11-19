package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"pooling-demo/pool"
	"sync"
	"time"
)

//resource
type DBConnection struct {
	ID int32
}

//Closer interface implementation
func (c *DBConnection) Close() error {
	println("Closing connection", c.ID)
	return nil
}

//resource factory
var idCounter int32

func DBConnectionFactory() (io.Closer, error) {
	idCounter++
	fmt.Println("DBConnectionFactory : Creating resource # ", idCounter)
	return &DBConnection{idCounter}, nil
}

func main() {
	p, err := pool.New(DBConnectionFactory /* factory */, 3 /* size */)
	if err != nil {
		log.Fatalln(err)
	}

	wg := &sync.WaitGroup{}
	clientCount := 10
	wg.Add(clientCount)
	for client := 0; client < clientCount; client++ {
		go func(client int) {
			doWork(client, p)
			wg.Done()
		}(client)
	}
	wg.Wait()

	fmt.Println("Second batch of operations")
	wg.Add(3)
	for idx := 10; idx < 13; idx++ {
		go func(client int) {
			doWork(client, p)
			wg.Done()
		}(idx)
	}
	wg.Wait()
	p.Close()
}

func doWork(id int, p *pool.Pool) {
	conn, err := p.Acquire()
	if err != nil {
		log.Fatalln(err)
	}
	defer p.Release(conn)

	//use the connection
	fmt.Println("Worker : ", id, " : Acquired : ", conn.(*DBConnection).ID)
	//do some work
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	fmt.Println("Worker Done : ", id, " : Release : ", conn.(*DBConnection).ID)
}

/*
   resource
		implements Closer interface (Close method)

	resource factor
		returns a new resource
*/
