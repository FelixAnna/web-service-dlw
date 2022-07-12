package repositories

import (
	"context"
	"log"
	"time"

	"github.com/FelixAnna/web-service-dlw/common/aws"
	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/entity"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoRepoSet = wire.NewSet(ProvideMongoQuestionRepo, wire.Bind(new(QuestionRepo), new(*MongoQuestionRepo)))

type MongoQuestionRepo struct {
	ClientOptions *options.ClientOptions
}

func ProvideMongoQuestionRepo(awsService *aws.AWSService) *MongoQuestionRepo {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)

	clientOptions := options.Client().
		ApplyURI(awsService.GetParameterByKey("mongo/connectionstring")).
		SetServerAPIOptions(serverAPIOptions)

	return &MongoQuestionRepo{
		ClientOptions: clientOptions,
	}
}

func (repo *MongoQuestionRepo) GetQuestion(id string) *entity.Questions {
	ctx, cancel, collection := getCollection(repo, "dlw_mathematicals", "questions")
	defer cancel()

	var result entity.Questions //= entity.Questions{}
	err := collection.FindOne(ctx, bson.D{{"_id", id}}).Decode(&result)
	if err != nil {
		log.Println(err)
		return nil
	}

	return &result
}

func (repo *MongoQuestionRepo) SaveQuestions(questions *entity.Questions) error {
	ctx, cancel, collection := getCollection(repo, "dlw_mathematicals", "questions")
	defer cancel()

	_, err := collection.InsertOne(ctx, *questions)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (repo *MongoQuestionRepo) SaveAnswers(answers *entity.Answers) error {
	ctx, cancel, collection := getCollection(repo, "dlw_mathematicals", "answers")
	defer cancel()

	_, err := collection.InsertOne(ctx, *answers)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func getCollection(repo *MongoQuestionRepo, dbName, collectionName string) (context.Context, context.CancelFunc, *mongo.Collection) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, repo.ClientOptions)
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database(dbName).Collection(collectionName)
	return ctx, cancel, collection
}
