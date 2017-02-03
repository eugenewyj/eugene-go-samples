package main

import (
	"sync"
	"fmt"
	"runtime"
)

var (
	counter int
	wg sync.WaitGroup
	mutex sync.Mutex
)

func main() {
	runtime.GOMAXPROCS(1)
	
	wg.Add(2)

	go incCounter(1)
	go incCounter(2)

	wg.Wait()
	fmt.Printf("Final Counter: %d\n", counter)
}

func incCounter(id int) {
	defer wg.Done()
	fmt.Printf("task %d work\n", id)
	for count := 0; count < 2; count++ {
		mutex.Lock()
		{
			fmt.Printf("task %d loop\n", id)
			value := counter

			runtime.Gosched()

			value ++

			counter = value
		}
		mutex.Unlock()
	}
	fmt.Printf("task %d end\n", id)
}