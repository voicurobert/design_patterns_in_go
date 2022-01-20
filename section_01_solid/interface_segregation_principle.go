package solid

import "fmt"

type Document struct {
}

type Machine interface {
	Print(d Document)
	Fax(d Document)
	Scan(d Document)
}

type MultiFunctionPrinter struct {
}

func (m *MultiFunctionPrinter) Print(d Document) {

}

func (m *MultiFunctionPrinter) Fax(d Document) {

}

func (m *MultiFunctionPrinter) Scan(d Document) {

}

type OldFashionedPrinter struct {
}

func (o *OldFashionedPrinter) Print(d Document) {
	// ok
}

func (o *OldFashionedPrinter) Fax(d Document) {
	// ??
	panic("operation not supported")
}

//Deprecated
func (o *OldFashionedPrinter) Scan(d Document) {
	// ??
	panic("operation not supported")
}

type Printer interface {
	Print(d Document)
}

type Faxing interface {
	Fax(d Document)
}

type Scanner interface {
	Scan(d Document)
}

type MyPrinter struct {
}

func (m *MyPrinter) Print(d Document) {

}

type Photocopier struct {
}

func (p *Photocopier) Print(d Document) {

}

func (p *Photocopier) Scan(d Document) {

}

type MultiFunctionDevice interface {
	Printer
	Scanner
	// Fax
}

type MultiFunctionMachine struct {
	printer Printer
	scanner Scanner
}

func (m MultiFunctionMachine) Print(d Document) {
	m.printer.Print(d)
}

func InterfaceSegregationPrinciple() {
	fmt.Println("Interface Segregation Principle...")

}
