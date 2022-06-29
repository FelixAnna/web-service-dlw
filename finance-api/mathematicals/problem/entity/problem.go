package entity

import "fmt"

type Problem struct {
	A int
	B int
	C int

	Op rune
}

func (p *Problem) PrintAll() string {
	return fmt.Sprintf("%v %c %v = %v", p.A, p.Op, p.B, p.C)
}

func (p *Problem) QuestFirst() string {
	return fmt.Sprintf("? %c %v = %v", p.Op, p.B, p.C)
}

func (p *Problem) QuestSecond() string {
	return fmt.Sprintf("%v %c ? = %v", p.A, p.Op, p.C)
}

func (p *Problem) QuestResult() string {
	return fmt.Sprintf("%v %c %v = ?", p.A, p.Op, p.B)
}
