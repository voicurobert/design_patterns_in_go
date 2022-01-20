package prototype

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

func (p *Person) DeepCopySerialization() *Person {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(p)

	fmt.Println(string(b.Bytes()))

	d := gob.NewDecoder(&b)

	result := Person{}
	_ = d.Decode(&result)

	return &result
}

func SerializationExample() {
	john := Person{
		"John",
		&Address{"123 London", "London", "UK"},
		[]string{"Chrish", "Mat"}}

	jane2 := john.DeepCopySerialization()
	jane2.name = "Jane2"
	jane2.Address.StreetAddress = "asdada"
	jane2.Friends = append(jane2.Friends, "Angela")
	fmt.Println(jane2, jane2.Address)
}
