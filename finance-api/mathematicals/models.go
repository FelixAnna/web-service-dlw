package mathematicals

import "math"

const (
	KindQuestFirst  int = 1
	KindQuestSecond int = 2
	KindQeustLast   int = 3
)

const (
	CategoryPlus int = iota
	CategoryMinus
)

type Range struct {
	Min, Max int
}

type Criteria struct {
	Min, Max int `binding:"-"`
	Quantity int `binding:"min=1,max=10000"`

	Range *Range `binding:"-"`

	//+, -
	Category int `binding:"min=0,max=1"`

	//first, second, last
	Kind int `binding:"min=1,max=3"`
}

func (s *Criteria) GetRange() (min, max int) {
	min, max = math.MinInt32, math.MaxInt32
	if s.Range == nil {
		return
	}

	if s.Range.Min > s.Range.Max {
		return s.Range.Max, s.Range.Min
	}

	return s.Range.Min, s.Range.Max
}

type QuestionModel struct {
	Question string
	Answer   int

	Kind     int
	FullText string
}
