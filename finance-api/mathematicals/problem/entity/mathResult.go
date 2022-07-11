package entity

type QuestionResult struct {
	Index    int
	Question string
	Answer   string

	Category int
	Kind     int
	Type     int

	UserAnswer int
}

type MathResult struct {
	Id     string
	UserId string

	GroupId         string
	QuestionResults []QuestionResult
}
