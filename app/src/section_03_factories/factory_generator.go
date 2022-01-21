package factory

import "fmt"

type Employee struct {
	Name, Position string
	AnnualIncome   int
}

// functional approach

func NewEmployeeFactory(position string, annualIncome int) func(name string) *Employee {
	return func(name string) *Employee {
		return &Employee{name, position, annualIncome}
	}
}

// structural approach

type EmployeeFactory struct {
	Position     string
	AnnualIncome int
}

func (f *EmployeeFactory) Create(name string) *Employee {
	return &Employee{name, f.Position, f.AnnualIncome}
}

func NewEmployeeFactory2(position string, annualIncome int) *EmployeeFactory {
	return &EmployeeFactory{position, annualIncome}
}

func GeneratorExample() {
	developerFactory := NewEmployeeFactory("Developer", 20000)
	managerFactory := NewEmployeeFactory("Manager", 80000)

	dev := developerFactory("Robert")
	man := managerFactory("Laur")

	fmt.Println()
	fmt.Println(dev)
	fmt.Println(man)

	bossFactory := NewEmployeeFactory2("CEO", 100000)
	boss := bossFactory.Create("Sam")
	fmt.Println(boss)
}
