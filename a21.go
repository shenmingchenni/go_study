package main

import "fmt"

func myfunc(Handleeverything interface{}) {
	value, ok := Handleeverything.(int)
	if ok {
		fmt.Printf("收到现金消费:%d元", value)
		return
	}
	dog, ok := Handleeverything.(Dog)
	if ok {
		fmt.Printf("欢迎新狗狗:%s,它是可爱的:%s", dog.Name, dog.Breed)
		return
	}
	food, ok := Handleeverything.(Food)
	if ok {
		fmt.Printf("入库%s牌狗粮,重量%dkg", food.Brand, food.Weight)
		return
	}
	member, ok := Handleeverything.(Member)
	if ok {
		fmt.Printf("会员%s到店,账户余额%f元", member.Name, member.Balance)
		return
	}
	fmt.Println("警报,收到不名物体")
}

type Dog struct {
	Name  string
	Breed string
}

type Food struct {
	Brand  string
	Weight int
}

type Member struct {
	Name    string
	Balance float64
}

func main() {
	dog := Dog{Name: "旺财", Breed: "金毛"}
	food := Food{Brand: "大大", Weight: 3}
	member := Member{Name: "张三", Balance: 12.3}
	myfunc(dog)
	myfunc(food)
	myfunc(member)
}
