package factory

import "fmt"

type Person struct {
	Name     string
	Age      int
	EyeCount int
}

func NewPerson(name string, age int) *Person {
	if age < 16 {
		// validation
	}

	return &Person{
		Name:     name,
		Age:      age,
		EyeCount: 2,
	}
}

func FunctionExample() {
	p := NewPerson("John", 22)
	fmt.Println(p)
}
