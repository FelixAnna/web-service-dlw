package entity

import "fmt"

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
