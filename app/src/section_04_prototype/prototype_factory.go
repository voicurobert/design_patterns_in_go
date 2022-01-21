package prototype

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Address2 struct {
	Suite               int
	StreetAddress, City string
}

type Employee struct {
	Name   string
	Office Address2
}

func (p *Employee) DeepCopy() *Employee {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(p)

	// peek into structure
	fmt.Println(string(b.Bytes()))

	d := gob.NewDecoder(&b)
	result := Employee{}
	_ = d.Decode(&result)
	return &result
}

var MainOffice = Employee{
	"",
	Address2{0, "123 East Dr", "London"},
}

var AuxOffice = Employee{
	"",
	Address2{0, "66 West Dr", "London"},
}

func newEmployee(proto *Employee, name string, suite int) *Employee {
	result := proto.DeepCopy()
	result.Name = name
	result.Office.Suite = suite

	return result
}

func NewMainOfficeEmployee(name string, suite int) *Employee {
	return newEmployee(&MainOffice, name, suite)
}

func NewAuxEmployee(name string, suite int) *Employee {
	return newEmployee(&AuxOffice, name, suite)
}

func FactoryExample() {
	john := NewMainOfficeEmployee("John", 44)

	jane := NewAuxEmployee("Jane", 100)

	fmt.Println(john, john.Office)
	fmt.Println(jane, jane.Office)

}
