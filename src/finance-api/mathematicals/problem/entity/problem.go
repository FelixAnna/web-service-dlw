package entity

import "fmt"

type Problem[number Number] struct {
	A number
	B number
	C number

	Op rune
}

func (p *Problem[number]) IndenticalString() string {
	a, b, c := p.A, p.B, p.C
	if p.Op == '+' || p.Op == '*' {
		if a > b {
			a, b = b, a
		}
	}

	return fmt.Sprintf("%v %c %v = %v", a, p.Op, b, c)
}
