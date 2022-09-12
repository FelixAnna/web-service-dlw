package format

import (
	"fmt"

	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/entity"
)

type PlainApplication struct {
	*entity.Problem
	Template string
	Ops      []string
}

func (p *PlainApplication) String() string {
	// 1. 比65多5的数字是()  -> a + b = c
	// 2. 比65少5的数字是（）-> a - b = c
	return fmt.Sprintf(p.Template, p.A, p.getOp(), p.B, p.C)
}

func (p *PlainApplication) QuestFirst() string {
	return fmt.Sprintf(p.Template, placeHolder, p.getOp(), p.B, p.C)
}

func (p *PlainApplication) QuestSecond() string {
	return fmt.Sprintf(p.Template, p.A, p.getOp(), placeHolder, p.C)
}

func (p *PlainApplication) QuestResult() string {
	return fmt.Sprintf(p.Template, p.A, p.getOp(), p.B, placeHolder)
}

//"比%v%s%v的数是%s"
func (p *PlainApplication) getOp() string {
	if p.Op == '+' {
		return p.Ops[0]
	} else {
		return p.Ops[1]
	}
}
