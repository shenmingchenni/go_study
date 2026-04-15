package main

import "fmt"

func main() {
	c := make(chan string, 5)
	go func() {
		for i := 0; i < 6; i++ {
			abc := fmt.Sprintf("订单:%d\n", i)
			c <- abc
			fmt.Printf("正在传送:%s\n", abc)
		}
	}()
	for i := 0; i < 6; i++ {
		num := <-c
		fmt.Printf("num为:%s\n", num)
	}
	fmt.Println("结束")
}
