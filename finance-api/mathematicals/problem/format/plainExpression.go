package format

import (
	"fmt"

	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/entity"
)

const placeHolder = "(  )"

type PlainExpression struct {
	*entity.Problem
}

func (p *PlainExpression) String() string {
	return fmt.Sprintf("%v %s %v = %v", p.A, p.getOp(), p.B, p.C)
}

func (p *PlainExpression) QuestFirst() string {
	return fmt.Sprintf("%s %s %v = %v", placeHolder, p.getOp(), p.B, p.C)
}

func (p *PlainExpression) QuestSecond() string {
	return fmt.Sprintf("%v %s %s = %v", p.A, p.getOp(), placeHolder, p.C)
}

func (p *PlainExpression) QuestResult() string {
	return fmt.Sprintf("%v %s %v = %s", p.A, p.getOp(), p.B, placeHolder)
}

func (p *PlainExpression) getOp() string {
	if p.Op == '+' {
		return "+"
	} else {
		return "-"
	}
}
