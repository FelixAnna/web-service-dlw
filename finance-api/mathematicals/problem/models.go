package problem

import (
	"math"
)

const (
	KindQuestFirst  int = 1
	KindQuestSecond int = 2
	KindQeustLast   int = 3
)

const (
	CategoryPlus int = iota
	CategoryMinus
)

const (
	TypePlainExpression int = iota
	TypePlainApplication
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

	//Expression, plainApplication, ...
	Type int `binding:"min=0,max=255"`
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

type QuestionResponse struct {
	Questions  []QuestionModel
	QuestionId string
}

type QuestionModel struct {
	Question string
	Answer   int

	Category int
	Kind     int
	Type     int
}

type QuestionFeedModel struct {
	Questions []string
	Answers   []string
	FullText  []string
}

type QuestionAnswerItem struct {
	Index    int    `json:"Index" binding:"-"`
	Question string `json:"Question" binding:"required"`
	Answer   string `json:"Answer" binding:"required"`

	Category int `json:"Category" binding:"-"`
	Kind     int `json:"Kind" binding:"required"`
	Type     int `json:"Type" binding:"-"`

	UserAnswer string `json:"UserAnswer" binding:"-"`
}

type SaveAnswersRequest struct {
	Results    []QuestionAnswerItem `json:"Results,omitempty" binding:"required,dive,required"`
	QuestionId string               `json:"QuestionId" binding:"required"`
	Score      float32              `json:"Score" binding:"-"`
}
