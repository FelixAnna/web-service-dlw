package entity

type QuestionItem struct {
	Index    int
	Question string
	Answer   string

	Category int
	Kind     int
	Type     int
}

type Questions struct {
	Id string `bson:"_id"`

	Questions []QuestionItem

	CreatedTime int64 `bson:"created_time"`
}

type AnswerItem struct {
	Index  int
	Answer string
}

type Answers struct {
	Id string `bson:"_id"`

	UserId     string `bson:"user_id"`
	QuestionId string `bson:"question_id"`
	Score      float32
	Answers    []AnswerItem

	CreatedTime int64 `bson:"created_time"`
}
