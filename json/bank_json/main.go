package main

import (
	"encoding/json"
	"fmt"
)

type Transaction struct {
	FromAccount string `json:"fromaccount"`
	ToAccount   string `json:"toaccount"`
	Amount      int    `json:"amount"`
}

func main() {
	bInput := `{"fromaccount":"hzy","toaccount":"ymj","amount":999999}`
	var a Transaction
	err := json.Unmarshal([]byte(bInput), &a)
	if err != nil {
		fmt.Print("不合法")
		return
	}
	if a.Amount <= 0 {
		fmt.Println("该转账不合法")
		return
	}
	a.Amount = 888
	cstr, err := json.Marshal(a)
	if err != nil {
		fmt.Println("不合法")
	}
	fmt.Printf("最终转账结果为:%s", cstr)
}
