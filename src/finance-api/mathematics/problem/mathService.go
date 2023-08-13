package problem

import (
	"fmt"
	"sync"
	"time"

	"github.com/FelixAnna/web-service-dlw/common/snowflake"
	"github.com/FelixAnna/web-service-dlw/finance-api/mathematics/problem/entity"
	"github.com/FelixAnna/web-service-dlw/finance-api/mathematics/problem/format"
	"github.com/FelixAnna/web-service-dlw/finance-api/mathematics/problem/repositories"
)

type MathService struct {
	genService *TwoGenerationService
	qaService  repositories.QuestionRepo
}

func NewMathService(genService *TwoGenerationService, qaService repositories.QuestionRepo) *MathService {
	return &MathService{
		genService: genService,
		qaService:  qaService,
	}
}

func (service *MathService) GenerateProblems(criteria ...Criteria) *QuestionResponse {
	var results []QuestionModel
	for _, criteria := range criteria {
		cr := criteria
		problems := GetResponse(service.genService.GenerateProblems(&cr), &cr)
		results = append(results, problems...)
	}

	return &QuestionResponse{
		Questions:  results,
		QuestionId: snowflake.GenerateSnowflake(),
	}
}

func (service *MathService) SaveResults(request *SaveAnswersRequest, userId string) error {
	err := ensureSaveQuestions(service, request)
	if err != nil {
		return err
	}

	err = saveAnswers(userId, request, service)
	return err
}

func saveAnswers(userId string, request *SaveAnswersRequest, service *MathService) error {
	answers := entity.Answers{
		Id: snowflake.GenerateSnowflake(),

		UserId:      userId,
		QuestionId:  request.QuestionId,
		Score:       request.Score,
		Answers:     []entity.AnswerItem{},
		CreatedTime: time.Now().UTC().Unix(),
	}

	for _, val := range request.Results {
		if val.UserAnswer == "" {
			continue
		}

		answers.Answers = append(answers.Answers, entity.AnswerItem{
			Answer: val.UserAnswer,
			Index:  val.Index,
		})
	}
	err := service.qaService.SaveAnswers(&answers)
	return err
}

func ensureSaveQuestions(service *MathService, request *SaveAnswersRequest) error {
	question := service.qaService.GetQuestion(request.QuestionId)
	if question == nil {
		questions := entity.Questions{
			Id:          request.QuestionId,
			Questions:   []entity.QuestionItem{},
			CreatedTime: time.Now().UTC().Unix(),
		}
		for _, val := range request.Results {
			question := entity.QuestionItem{
				Index:    val.Index,
				Question: val.Question,
				Answer:   val.Answer,

				Category: val.Category,
				Kind:     val.Kind,
				Type:     val.Type,
			}

			questions.Questions = append(questions.Questions, question)
		}

		return service.qaService.SaveQuestions(&questions)
	}

	return nil
}

func (service *MathService) GenerateFeeds(criteria ...Criteria) *QuestionFeedModel {
	ch := make(chan []string)
	wg := &sync.WaitGroup{}
	wg.Add(len(criteria))
	for _, criteria := range criteria {
		cr := criteria
		go GetResponseFeed(service.genService.GenerateProblems(&cr), &cr, wg, ch)
	}

	result, wg2 := processingResults(ch)

	//wait for generating
	wg.Wait()
	close(ch)

	//wait for result processing
	wg2.Wait()

	return result
}

func GetResponse(results []entity.Problem, criteria *Criteria) []QuestionModel {
	questions := []QuestionModel{}

	for _, pb := range results {
		expression := getFormatInterface(&pb, criteria)
		question, answer := getDisplayQandA(criteria, expression, pb)

		model := QuestionModel{
			//FullText: expression.String(),
			Kind:     criteria.Kind,
			Category: criteria.Category,
			Type:     criteria.Type,

			Question: question,
			Answer:   answer,
		}

		questions = append(questions, model)
	}

	return questions
}

func GetResponseFeed(results []entity.Problem, criteria *Criteria, wg *sync.WaitGroup, ch chan<- []string) {
	defer wg.Done()

	for _, pb := range results {
		expression := getFormatInterface(&pb, criteria)
		question, answer := getDisplayQandA(criteria, expression, pb)
		ch <- []string{question, fmt.Sprintf("%v", answer), expression.String()}
	}
}

func getDisplayQandA(criteria *Criteria, expression format.FormatInterface, pb entity.Problem) (string, int) {
	var question string
	var answer int
	switch criteria.Kind {
	case KindQuestFirst:
		question = expression.QuestFirst()
		answer = pb.A
	case KindQuestSecond:
		question = expression.QuestSecond()
		answer = pb.B
	case KindQeustLast:
		question = expression.QuestResult()
		answer = pb.C
	}
	return question, answer
}

func formart(input interface{}, idx int) string {
	return fmt.Sprintf("%v. %v", idx, input)
}

func processingResults(ch chan []string) (*QuestionFeedModel, *sync.WaitGroup) {
	var result = &QuestionFeedModel{}
	wg2 := &sync.WaitGroup{}
	wg2.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg2.Done()
		idx := 0
		for vals := range ch {
			idx++
			result.Questions = append(result.Questions, formart(vals[0], idx))
			result.Answers = append(result.Answers, formart(vals[1], idx))
			result.FullText = append(result.FullText, formart(vals[2], idx))
		}
	}(wg2)
	return result, wg2
}

func getFormatInterface(pb *entity.Problem, criteria *Criteria) format.FormatInterface {
	var expression format.FormatInterface
	switch criteria.Type {
	case TypePlainExpression:
		expression = &format.PlainExpression{
			Problem: pb,
		}
	case TypePlainApplication:
		expression = &format.PlainApplication{
			Problem:  pb,
			Template: []string{"比%v%s%v的数是%v", "%v的%v%s是%v"},
			Ops:      []string{"多", "少", "倍", "分之一"},
		}
	case TypeAppleApplication:
		expression = &format.PlainApplication{
			Problem:  pb,
			Template: []string{"小明有%v个苹果，小红比小明%s%v个，小红有%v个苹果？", "小明有%v个苹果，小红的苹果是小明的%v%s，小红有%v个苹果？"},
			Ops:      []string{"多", "少", "倍", "分之一"},
		}
	case TypeTemplateApplication:
		expression = &format.PlainApplication{
			Problem:  pb,
			Template: []string{"哥哥身高%v厘米，妹妹比哥哥%s%v厘米，妹妹身高%v厘米？", "哥哥身高%v厘米，妹妹的身高是哥哥的%v%s，妹妹身高%v厘米？"},
			Ops:      []string{"高", "矮", "倍", "分之一"},
		}
	}

	return expression
}
