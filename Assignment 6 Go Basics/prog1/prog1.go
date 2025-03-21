package main

import (
	"fmt"
	"time"
)

func power_h(n, p, res int) {
	if p == 1 {
		fmt.Println("Result:", res)
		return
	}
	go power_h(n, p-1, res*n)
}

func power(n, p int) {
	go power_h(n, p, n)
}

func main() {
	n, p := 2, 5
	power(n, p)
	time.Sleep(2 * time.Second)
}
