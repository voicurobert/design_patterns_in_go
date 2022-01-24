package observer

import (
	"container/list"
	"fmt"
)

type Person2 struct {
	Observable
	age int
}

func NewPerson2(age int) *Person2 {
	return &Person2{
		Observable: Observable{new(list.List)},
		age:        age,
	}
}

type PropertyChange struct {
	Name  string // "Age"
	Value interface{}
}

func (p *Person2) Age() int {
	return p.age
}

func (p *Person2) SetAge(age int) {
	if age == p.age {
		return
	}
	p.age = age
	p.Fire(PropertyChange{"Age", p.age})
}

type TrafficManagement struct {
	o Observable
}

func (t TrafficManagement) Notify(data interface{}) {
	if pc, ok := data.(PropertyChange); ok {
		if pc.Value.(int) >= 18 {
			fmt.Println("You can drive now!")
			t.o.Unsubscribe(t)
		}
	}
}

func PropertyObserversExample() {
	p := NewPerson2(15)
	t := &TrafficManagement{p.Observable}

	p.Subscribe(t)

	for i := 16; i <= 20; i++ {
		fmt.Println("Setting the age to", i)
		p.SetAge(i)

	}

}
