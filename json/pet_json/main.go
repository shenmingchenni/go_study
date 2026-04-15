package main

import (
	"encoding/json"
	"fmt"
)

type Order struct {
	Name   string   `json:"name"`
	Price  float64  `json:"price"`
	Items  []string `json:"items"`
	Status string   `json:"status"`
}

func main() {
	jsonInput := `{"name":"旺财","price":125.5,"items":["猫粮","罐头"],"status":"pending"}`
	var o Order
	err := json.Unmarshal([]byte(jsonInput), &o)
	if err != nil {
		fmt.Println("错误", err)
		return
	}
	if o.Price > 100 {
		fmt.Println("赠送猫条")
	}
	o.Status = "shipped"
	jsonstr, err := json.Marshal(o)
	if err != nil {
		fmt.Println("错误", err)
	}
	fmt.Printf("最终发货:%s", jsonstr)
}
