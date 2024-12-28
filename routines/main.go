package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(200 * time.Millisecond)
	go fib(40)
	time.Sleep(10 * time.Second)
	fmt.Printf("All done")
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) {
	for i := 1; i <= x; i++ {
		fmt.Printf("fib(%d)=%d\n", i, fib0(i))
	}
}

func fib0(x int) int {
	if x < 2 {
		return x
	}
	return fib0(x-1) + fib0(x-2)
}
