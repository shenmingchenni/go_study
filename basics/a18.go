package main

import "fmt"

type Hero struct {
	name  string
	ad    int
	level int
}

func (this *Hero) show() {
	fmt.Println("name=", this.name)
	fmt.Println("ad=", this.ad)
	fmt.Println("level=", this.level)
}

func (this *Hero) setname(newname string) {
	this.name = newname
}

func (this *Hero) getname() string {
	return this.name
}

func main() {
	hero := Hero{name: "hzy", ad: 1, level: 2}
	hero.show()
	hero.setname("lisi")
	hero.show()

}
