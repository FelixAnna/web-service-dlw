package entity

import "fmt"

const placeHolder = "(  )"

type Problem struct {
	A int
	B int
	C int

	Op rune
}

func (p *Problem) IndenticalString() string {
	a, b, c := p.A, p.B, p.C
	if p.Op == '+' /*|| p.Op == '*' */ {
		if a > b {
			a, b = b, a
		}
	}

	return fmt.Sprintf("%v %c %v = %v", a, p.Op, b, c)
}

func (p *Problem) String() string {
	return fmt.Sprintf("%v %c %v = %v", p.A, p.Op, p.B, p.C)
}

func (p *Problem) QuestFirst() string {
	return fmt.Sprintf("%s %c %v = %v", placeHolder, p.Op, p.B, p.C)
}

func (p *Problem) QuestSecond() string {
	return fmt.Sprintf("%v %c %s = %v", p.A, p.Op, placeHolder, p.C)
}

func (p *Problem) QuestResult() string {
	return fmt.Sprintf("%v %c %v = %s", p.A, p.Op, p.B, placeHolder)
}
