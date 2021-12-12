package visitor

import (
	"fmt"
	"strings"
)

type ExpressionReflective interface {
	// nothing here!
}

type DoubleExpressionReflective struct {
	value float64
}

type AdditionExpressionReflective struct {
	left, right ExpressionReflective
}

func Print(e ExpressionReflective, sb *strings.Builder) {
	if de, ok := e.(*DoubleExpressionReflective); ok {
		sb.WriteString(fmt.Sprintf("%g", de.value))
	} else if ae, ok := e.(*AdditionExpressionReflective); ok {
		sb.WriteString("(")
		Print(ae.left, sb)
		sb.WriteString("+")
		Print(ae.right, sb)
		sb.WriteString(")")
	}

	// breaks OCP
	// will work incorrectly on missing case
}

func StartReflectiveVistor() {
	// 1+(2+3)
	e := &AdditionExpressionReflective{
		&DoubleExpressionReflective{1},
		&AdditionExpressionReflective{
			left:  &DoubleExpressionReflective{2},
			right: &DoubleExpressionReflective{3},
		},
	}
	sb := strings.Builder{}
	Print(e, &sb)
	fmt.Println(sb.String())
}
