package format

import (
	"fmt"

	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/entity"
)

type PlainExpression[number entity.Number] struct {
	*entity.Problem[number]
}

func (p *PlainExpression[number]) String() string {
	return fmt.Sprintf("%v %s %v = %v", p.A, p.getOp(), p.B, p.C)
}

func (p *PlainExpression[number]) QuestFirst() string {
	return fmt.Sprintf("%s %s %v = %v", placeHolder, p.getOp(), p.B, p.C)
}

func (p *PlainExpression[number]) QuestSecond() string {
	return fmt.Sprintf("%v %s %s = %v", p.A, p.getOp(), placeHolder, p.C)
}

func (p *PlainExpression[number]) QuestResult() string {
	return fmt.Sprintf("%v %s %v = %s", p.A, p.getOp(), p.B, placeHolder)
}

func (p *PlainExpression[number]) getOp() string {
	if p.Op == '+' {
		return "+"
	} else if p.Op == '-' {
		return "-"
	} else if p.Op == '*' {
		return "*"
	} else if p.Op == '/' {
		return "/"
	} else {
		return "?" //not supported
	}
}
