package visitor

import (
	"fmt"
	"strings"
)

type ExpressionVisitor interface {
	VisitDoubleExpression(de *DoubleExpressionClassic)
	VisitAdditionExpression(ae *AdditionExpressionClassic)
}

type ExpressionClassic interface {
	Accept(ev ExpressionVisitor)
}

type DoubleExpressionClassic struct {
	value float64
}

func (d *DoubleExpressionClassic) Accept(ev ExpressionVisitor) {
	ev.VisitDoubleExpression(d)
}

type AdditionExpressionClassic struct {
	left, right ExpressionClassic
}

func (a *AdditionExpressionClassic) Accept(ev ExpressionVisitor) {
	ev.VisitAdditionExpression(a)
}

type ExpressionPrinter struct {
	sb strings.Builder
}

func (e *ExpressionPrinter) VisitDoubleExpression(de *DoubleExpressionClassic) {
	e.sb.WriteString(fmt.Sprintf("%g", de.value))
}

func (e *ExpressionPrinter) VisitAdditionExpression(ae *AdditionExpressionClassic) {
	e.sb.WriteString("(")
	ae.left.Accept(e)
	e.sb.WriteString("+")
	ae.right.Accept(e)
	e.sb.WriteString(")")
}

func NewExpressionPrinter() *ExpressionPrinter {
	return &ExpressionPrinter{strings.Builder{}}
}

func (e *ExpressionPrinter) String() string {
	return e.sb.String()
}

func StartClassicVisitor() {
	// 1+(2+3)
	e := &AdditionExpressionClassic{
		&DoubleExpressionClassic{1},
		&AdditionExpressionClassic{
			left:  &DoubleExpressionClassic{2},
			right: &DoubleExpressionClassic{3},
		},
	}
	ep := NewExpressionPrinter()
	ep.VisitAdditionExpression(e)
	fmt.Println(ep.String())
}
