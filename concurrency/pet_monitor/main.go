package main

import (
	"fmt"
	"time"
)

type Order struct {
	Id     int
	Status string
}

func monitor(orderChan chan int, statusChan chan string, stopChan chan struct{}) {
	for {
		select {
		case data := <-orderChan:
			fmt.Printf("收到订单:%d\n", data)
		case status := <-statusChan:
			fmt.Printf("系统状态:%s\n", status)
		case stop := <-stopChan:
			fmt.Printf("结束:%v\n", stop)
			return
		}
	}
}
func main() {
	orderChan := make(chan int)
	statusChan := make(chan string)
	stopChan := make(chan struct{})
	go func() {
		for i := 0; i < 10; i++ {
			orderChan <- i
			statusChan <- "订单产生在"
			time.Sleep(500 * time.Millisecond)

		}
		stopChan <- struct{}{}
	}()
	go func() {
		for i := 0; ; i++ {
			time.Sleep(time.Second)
			statusChan <- "系统运行正常"
		}
	}()
	monitor(orderChan, statusChan, stopChan)
}
