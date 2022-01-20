package factory

import "fmt"

type person struct {
	name string
	age  int
}

type tiredPerson struct {
	name string
	age  int
}

type PersonInterface interface {
	SayHello()
}

func (p *person) SayHello() {
	fmt.Printf("Hi, my name is %s, I am %d years old\n", p.name, p.age)
}

func (p *tiredPerson) SayHello() {
	fmt.Printf("Sorry, I'm too tired")
}

func NewPersonInterface(name string, age int) PersonInterface {
	if age > 100 {
		return &tiredPerson{name, age}
	}
	return &person{name, age}
}

func InterfaceExample() {
	p := NewPersonInterface("James", 134)
	p.SayHello()
}
