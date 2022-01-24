package iterator

import "fmt"

type Person struct {
	FirstName, MiddleName, LastName string
}

func (p *Person) Names() [3]string {
	return [3]string{p.FirstName, p.MiddleName, p.LastName}
}

func (p *Person) NamesGenerator() <-chan string {
	out := make(chan string)

	go func() {
		defer close(out)
		out <- p.FirstName
		if len(p.MiddleName) > 0 {
			out <- p.MiddleName
		}
		out <- p.LastName
	}()

	return out
}

type PersonNameIterator struct {
	person  *Person
	current int
}

func NewPersonIterator(p *Person) *PersonNameIterator {
	return &PersonNameIterator{p, -1}
}

func (p *PersonNameIterator) MoveNext() bool {
	p.current++
	return p.current < 3
}

func (p *PersonNameIterator) Value() string {
	switch p.current {
	case 0:
		return p.person.FirstName
	case 1:
		return p.person.MiddleName
	case 2:
		return p.person.LastName
	default:
		panic("we should not be here")
	}
}

func IterationExample() {
	p := Person{"Alexander", "Graham", "Bell"}
	for name := range p.NamesGenerator() {
		fmt.Println(name)
	}

	for it := NewPersonIterator(&p); it.MoveNext(); {
		fmt.Println(it.Value())
	}
}
