package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 3)
	go func() {
		defer fmt.Println("Goroutine运行结束")
		fmt.Println("Goroutine运行中")
		for i := 0; i < 4; i++ {
			c <- i
			fmt.Printf("子协程发送:%d\n", i)
		}
	}()
	time.Sleep(time.Second)
	for i := 0; i < 4; i++ {
		num := <-c
		fmt.Printf("num=:%d\n", num)
	}
	fmt.Println("运行结束")
}
