package entity

type QuestionItem struct {
	Index    int    `json:"index" binding:"required"`
	Question string `json:"question" binding:"required"`
	Answer   string `json:"answer" binding:"required"`

	Category int `json:"category" binding:"required"`
	Kind     int `json:"kind" binding:"required"`
	Type     int `json:"type" binding:"required"`
}

type Questions struct {
	Id string `json:"_id" binding:""`

	Questions []QuestionItem `json:"questions" binding:""`

	CreatedTime int64 `json:"created_time" binding:""`
}

type AnswerItem struct {
	Index  int    `json:"index" binding:"required"`
	Answer string `json:"answer" binding:"required"`
}

type Answers struct {
	Id string `json:"_id" binding:""`

	UserId     string       `json:"user_id" binding:"required"`
	QuestionId string       `json:"question_id" binding:"required"`
	Score      float32      `json:"score" binding:""`
	Answers    []AnswerItem `json:"answers" binding:"required"`

	CreatedTime int64 `json:"created_time" binding:"required"`
}
