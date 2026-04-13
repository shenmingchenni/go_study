package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		defer fmt.Println("goroutine运行结束")
		fmt.Println("goroutine运行中")
		c <- 777
	}()
	num := <-c
	fmt.Printf("num为:%d\n", num)
	fmt.Println("结束")
}
