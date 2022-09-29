package format

import (
	"fmt"

	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/entity"
)

type PlainApplication struct {
	*entity.Problem
	Template []string
	Ops      []string
}

func (p *PlainApplication) String() string {
	// 1. 比65多5的数字是()  -> a + b = c
	// 2. 比65少5的数字是（）-> a - b = c
	return fmt.Sprintf(p.getTemplate(), p.A, p.getOp(), p.B, p.C)
}

func (p *PlainApplication) QuestFirst() string {
	return p.Sprintf(placeHolder, p.B, p.C)
}

func (p *PlainApplication) QuestSecond() string {
	return p.Sprintf(p.A, placeHolder, p.C)
}

func (p *PlainApplication) QuestResult() string {
	return p.Sprintf(p.A, p.B, placeHolder)
}

func (p *PlainApplication) Sprintf(a, b, c interface{}) string {
	if p.Op == '+' || p.Op == '-' {
		return fmt.Sprintf(p.Template[0], a, p.getOp(), b, c)
	} else if p.Op == '*' || p.Op == '/' {
		return fmt.Sprintf(p.Template[1], a, b, p.getOp(), c)
	} else {
		return "?" //not supported
	}
}
func (p *PlainApplication) getTemplate() string {
	if p.Op == '+' || p.Op == '-' {
		return p.Template[0]
	} else if p.Op == '*' {
		return p.Template[1]
	} else {
		return "?" //not supported
	}
}

//"比%v%s%v的数是%s"
func (p *PlainApplication) getOp() string {
	if p.Op == '+' {
		return p.Ops[0]
	} else if p.Op == '-' {
		return p.Ops[1]
	} else if p.Op == '*' {
		return p.Ops[2]
	} else {
		return "?" //not supported
	}
}
