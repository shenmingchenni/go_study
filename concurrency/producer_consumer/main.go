package main

import (
	"encoding/json"
	"fmt"
)

type Order struct {
	Id      int     `json:"id"`
	PetName string  `json:"petname"`
	Price   float64 `json:"price"`
}

func main() {
	c := make(chan []byte, 3)
	go func() {
		for i := 0; i < 5; i++ {
			o := Order{Id: i, PetName: "小狗", Price: 12.3}
			data, err := json.Marshal(o)
			if err != nil {
				fmt.Println(" 序列化失败: %v\n", err)
				continue
			}
			c <- data
			fmt.Printf("已发送订单:%d\n", i)
		}
		close(c)
	}()
	for astr := range c {
		var res Order
		err := json.Unmarshal(astr, &res)
		if err != nil {
			fmt.Println("序列化失败: %v\n", err)
			continue
		}
		fmt.Printf("成功处理订单:Id:%d,PetName:%s\n", res.Id, res.PetName)
	}
	fmt.Println("结束")
}
