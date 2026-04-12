package main

import (
	"fmt"
	"time"
)

func newtask() {
	i := 0
	for {
		i++
		fmt.Printf("newtask:=%d\n", i)
		time.Sleep(2 * time.Second)
	}
}
func main() {
	go newtask()
	i := 0
	for {
		i++
		fmt.Printf("maintask:=%d\n", i)
		time.Sleep(2 * time.Second)
	}
}
