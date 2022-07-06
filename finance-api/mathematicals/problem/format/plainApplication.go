package format

import (
	"fmt"

	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/entity"
)

type PlainApplication struct {
	*entity.Problem
}

func (p *PlainApplication) String() string {
	// 1. 比65多5的数字是()  -> a + b = c
	// 2. 比65少5的数字是（）-> a - b = c
	return fmt.Sprintf("比%v%s%v的数是%v", p.A, p.getOp(), p.B, p.C)
}

func (p *PlainApplication) QuestFirst() string {
	return fmt.Sprintf("比%s%s%v的数是%v", placeHolder, p.getOp(), p.B, p.C)
}

func (p *PlainApplication) QuestSecond() string {
	return fmt.Sprintf("比%v%s%s的数是%v", p.A, p.getOp(), placeHolder, p.C)
}

func (p *PlainApplication) QuestResult() string {
	return fmt.Sprintf("比%v%s%v的数是%s", p.A, p.getOp(), p.B, placeHolder)
}

func (p *PlainApplication) getOp() string {
	if p.Op == '+' {
		return "多"
	} else {
		return "少"
	}
}
