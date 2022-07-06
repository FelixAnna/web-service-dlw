package format

import (
	"testing"

	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/entity"
	"github.com/stretchr/testify/assert"
)

var plusApp PlainApplication
var plusApp2 PlainApplication
var minusApp PlainApplication

func init() {
	plusApp = PlainApplication{
		&entity.Problem{
			A:  1,
			B:  2,
			C:  3,
			Op: '+',
		},
	}

	plusApp2 = PlainApplication{
		&entity.Problem{
			A:  2,
			B:  1,
			C:  3,
			Op: '+',
		},
	}

	minusApp = PlainApplication{
		&entity.Problem{
			A:  3,
			B:  2,
			C:  1,
			Op: '-',
		},
	}
}

func TestAppString(t *testing.T) {
	result := plusApp.String()
	assert.EqualValues(t, "比1多2的数是3", result)
}

func TestAppIdenticalString(t *testing.T) {
	result := plusApp.IndenticalString()
	result2 := plusApp2.IndenticalString()
	assert.EqualValues(t, result2, result)
}

func TestAppNotIdenticalString(t *testing.T) {
	result := plusApp.IndenticalString()
	result2 := minusApp.IndenticalString()
	assert.NotEqualValues(t, result2, result)
}

func TestAppPrintFirst(t *testing.T) {
	result := plusApp.QuestFirst()
	assert.EqualValues(t, "比(  )多2的数是3", result)
}

func TestAppPrintSecond(t *testing.T) {
	result := plusApp.QuestSecond()
	assert.EqualValues(t, "比1多(  )的数是3", result)
}

func TestAppPrintLast(t *testing.T) {
	result := plusApp.QuestResult()
	assert.EqualValues(t, "比1多2的数是(  )", result)
}

func TestAppPrintMinusLast(t *testing.T) {
	result := minusApp.QuestResult()
	assert.EqualValues(t, "比3少2的数是(  )", result)
}
