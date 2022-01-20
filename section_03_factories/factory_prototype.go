package factory

import "fmt"

const (
	Developer = iota
	Manager
)

func NewEmployee(role int) *Employee {
	switch role {
	case Developer:
		return &Employee{"", "Developer", 60000}
	case Manager:
		return &Employee{"", "Manager", 90000}
	default:
		panic("Unsupported role")
	}
}

func PrototypeExample() {
	m := NewEmployee(Manager)
	m.Name = "Sam"
	fmt.Println(m)
}
