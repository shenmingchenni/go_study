package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	task := []string{"清理猫砂", "遛狗归来", "检查库存"}
	for i := 0; i < len(task); i++ {
		wg.Add(1)
		go func(taskname string) {
			defer wg.Done()
			fmt.Printf("员工正在:%s\n", taskname)
			time.Sleep(time.Second)
			fmt.Printf("员工已经:%s\n", taskname)
		}(task[i])
	}
	fmt.Println("老板在等待")
	wg.Wait()
	fmt.Println("结束")
}
