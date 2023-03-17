package problem

import (
	"math"

	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/entity"
)

const (
	KindQuestFirst  int = 1
	KindQuestSecond int = 2
	KindQeustLast   int = 3
)

const (
	CategoryPlus int = iota
	CategoryMinus
	CategoryMultiply
	CategoryDivide
)

const (
	TypePlainExpression int = iota
	TypePlainApplication
	TypeAppleApplication
	TypeTemplateApplication
)

type Range struct {
	Min, Max int
}

type Criteria[number entity.Number] struct {
	Min, Max int `binding:"-"`
	Quantity int `binding:"min=1,max=10000"`

	Range *Range `binding:"-"`

	//+, -
	Category int `binding:"min=0,max=3"`

	//first, second, last
	Kind int `binding:"min=1,max=3"`

	//Expression, plainApplication, ...
	Type int `binding:"min=0,max=255"`
}

func (s *Criteria[number]) GetRange() (min, max number) {
	min, max = math.MinInt32, math.MaxInt32
	if s.Range == nil {
		return
	}

	if s.Range.Min > s.Range.Max {
		return (interface{}(s.Range.Max)).(number), (interface{}(s.Range.Min)).(number)
	}

	return (interface{}(s.Range.Min)).(number), (interface{}(s.Range.Max)).(number)
}

type QuestionResponse[number entity.Number] struct {
	Questions  []QuestionModel[number]
	QuestionId string
}

type QuestionModel[number entity.Number] struct {
	Question string
	Answer   number

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
