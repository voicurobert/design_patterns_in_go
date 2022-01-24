package visitor

import (
	"fmt"
	"strings"
)

type ExpressionVisitor interface {
	VisitDoubleExpression(e *DoubleExpression)
	VisitAdditionExpression(e *AdditionExpression)
}

type Expression interface {
	//Print(sb *strings.Builder)

	Accept(ev ExpressionVisitor)
}

type DoubleExpression struct {
	value float64
}

func (d *DoubleExpression) Print(sb *strings.Builder) {
	sb.WriteString(fmt.Sprintf("%g", d.value))
}

func (d *DoubleExpression) Accept(ev ExpressionVisitor) {
	ev.VisitDoubleExpression(d)
}

type AdditionExpression struct {
	left, right Expression
}

func (a *AdditionExpression) Accept(ev ExpressionVisitor) {
	ev.VisitAdditionExpression(a)
}

type ExpressionPrinter struct {
	sb strings.Builder
}

func (ep *ExpressionPrinter) VisitDoubleExpression(e *DoubleExpression) {
	ep.sb.WriteString(fmt.Sprintf("%g", e.value))
}

func (ep *ExpressionPrinter) VisitAdditionExpression(e *AdditionExpression) {
	ep.sb.WriteRune('(')
	e.left.Accept(ep)
	ep.sb.WriteRune('+')
	e.right.Accept(ep)
	ep.sb.WriteRune(')')
}

func (ep *ExpressionPrinter) String() string {
	return ep.sb.String()
}

func NewExpressionPrinter() *ExpressionPrinter {
	return &ExpressionPrinter{strings.Builder{}}
}

func (a AdditionExpression) Print(sb *strings.Builder) {
	sb.WriteRune('(')
	//	a.left.Print(sb)
	sb.WriteRune('+')
	//	a.right.Print(sb)
	sb.WriteRune(')')
}

//func Print(e Expression, sb *strings.Builder) {
//	if de, ok := e.(*DoubleExpression); ok {
//		sb.WriteString(fmt.Sprintf("%g", de.value))
//	} else if ae, ok := e.(*AdditionExpression); ok {
//		sb.WriteRune('(')
//		Print(ae.left, sb)
//		sb.WriteRune('+')
//		Print(ae.right, sb)
//		sb.WriteRune(')')
//	}
//}

func IntrusiveVisitorExample() {
	e := &AdditionExpression{
		left: &DoubleExpression{1},
		right: &AdditionExpression{
			left:  &DoubleExpression{2},
			right: &DoubleExpression{3},
		},
	}
	//sb := &strings.Builder{}
	//e.Print(sb)
	//Print(e, sb)
	ep := NewExpressionPrinter()
	e.Accept(ep)
	fmt.Println(ep.String())
}
