package main

import "fmt"

type Humen struct {
	name string
	sex  string
}

func (this *Humen) Eat() {
	fmt.Print("Eat()....")
}

func (this *Humen) Walk() {
	fmt.Println("Walk()....")
}

func main() {
	h := Humen{"lisi", "female"}
	h.Eat()
	h.Walk()
}
