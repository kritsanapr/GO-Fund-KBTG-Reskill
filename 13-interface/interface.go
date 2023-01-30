package main

import "fmt"

type  Phone interface {
	Call(number string)
}

type Home interface {
	Live(string string)
}

type Samsung struct {
	Name string
}

type Nockdown struct {
	Name string
}

func (s Samsung) Call(number string) {
	fmt.Println("Samsung call to", number)
}

func (h Nockdown) Live(string string) {
	fmt.Println("Nockdown live")
}

func Dial(p Phone, h Home) {
	p.Call("+1 123 456 789")
	h.Live("")
}

func main() {
	s := Samsung{Name: "Android"}
	h := Nockdown{Name: "Nockdown"}
	Dial(s,h)
}