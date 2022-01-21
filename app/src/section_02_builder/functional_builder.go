package builder

import "fmt"

type Person2 struct {
	name, position string
}

type personMod func(person *Person2)

type PersonBuilder2 struct {
	actions []personMod
}

func (b *PersonBuilder2) Called(name string) *PersonBuilder2 {
	b.actions = append(b.actions, func(person *Person2) {
		person.name = name
	})
	return b
}

func (b *PersonBuilder2) Build() *Person2 {
	p := Person2{}
	for _, a := range b.actions {
		a(&p)
	}
	return &p
}

func (b *PersonBuilder2) WorksAsA(position string) *PersonBuilder2 {
	b.actions = append(b.actions, func(person *Person2) {
		person.position = position
	})
	return b
}

func FunctionalExample() {
	b := PersonBuilder2{}
	p := b.Called("Dmitri").WorksAsA("Developer").Build()
	fmt.Println(*p)

}
