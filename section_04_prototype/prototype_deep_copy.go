package prototype

import "fmt"

type Address struct {
	StreetAddress, City, Country string
}

func (a *Address) DeepCopy() *Address {
	return &Address{a.StreetAddress, a.City, a.Country}
}

type Person struct {
	name    string
	Address *Address
	Friends []string
}

func (p *Person) DeepCopy() *Person {
	q := *p
	q.Address = p.Address.DeepCopy()
	copy(q.Friends, p.Friends)
	return &q
}

func DeepCopyExample() {
	john := Person{
		"John",
		&Address{"123 London", "London", "UK"},
		[]string{"Chrish", "Mat"}}

	//jane := john
	//
	//jane.name = "Jane" // ok
	//jane.Address.StreetAddress = "321 Baker st"

	// deep copy
	jane := john
	jane.Address = &Address{john.Address.StreetAddress, john.Address.City, john.Address.Country}
	jane.Address.StreetAddress = "321 Baker st"

	fmt.Println(john, john.Address)
	fmt.Println(jane, jane.Address)

	jane2 := john.DeepCopy()
	jane2.name = "Jane2"
	jane2.Address.StreetAddress = "asdada"
	jane2.Friends = append(jane2.Friends, "Angela")
	fmt.Println(jane2, jane2.Address)
}
