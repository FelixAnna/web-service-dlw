package format

const placeHolder = "(___)"

type FormatInterface interface {
	String() string
	QuestFirst() string
	QuestSecond() string
	QuestResult() string
	IndenticalString() string
}
