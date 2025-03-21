package main

import (
	"fmt"
	"sync"
	"time"
)

var res int = 1
var mutex sync.Mutex

func power(n, p int) {
	var wg sync.WaitGroup

	for i := 0; i < p; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mutex.Lock()
			res *= n
			mutex.Unlock()
		}()
	}

	wg.Wait()
}

func main() {
	n := 2
	p := 6
	power(n, p)
	time.Sleep(2 * time.Second)
	fmt.Printf("Result: %d\n", res)
}
