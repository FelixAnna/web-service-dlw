package repository

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/FelixAnna/web-service-dlw/common/aws"
	"github.com/FelixAnna/web-service-dlw/memo-api/memo/entity"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var MongoRepoSet = wire.NewSet(ProvideMemoRepoMongoDB, wire.Bind(new(MemoRepo), new(*MemoRepoMongoDB)))

type MemoRepoMongoDB struct {
	ClientOptions *options.ClientOptions
	collection    string
	dbMame        string
}

func ProvideMemoRepoMongoDB(awsService *aws.AWSService) *MemoRepoMongoDB {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)

	clientOptions := options.Client().
		ApplyURI(awsService.GetParameterByKey("mongo/connectionstring")).
		SetServerAPIOptions(serverAPIOptions)

	return &MemoRepoMongoDB{
		ClientOptions: clientOptions,
		dbMame:        "dlw_memo",
		collection:    "memos",
	}
}

func (repo *MemoRepoMongoDB) Add(memo *entity.Memo) (string, error) {
	randId := fmt.Sprintf("%d%03d", time.Now().Unix(), rand.Intn(1000))
	memo.Id = randId
	memo.CreateTime = strconv.FormatInt(time.Now().UTC().Unix(), 10)

	ctx, cancel, collection := getCollection(repo, repo.dbMame, repo.collection)
	defer cancel()

	_, err := collection.InsertOne(ctx, *memo)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return memo.Id, nil
}

func (repo *MemoRepoMongoDB) GetById(id string) (*entity.Memo, error) {
	ctx, cancel, collection := getCollection(repo, repo.dbMame, repo.collection)
	defer cancel()

	var result entity.Memo //= entity.Questions{}
	err := collection.FindOne(ctx, bson.D{{Key: "_id", Value: id}}).Decode(&result)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &result, nil
}

func (repo *MemoRepoMongoDB) GetByUserId(userId string) ([]entity.Memo, error) {
	ctx, cancel, collection := getCollection(repo, repo.dbMame, repo.collection)
	defer cancel()

	var results []entity.Memo
	cursor, err := collection.Find(ctx, bson.D{{Key: "userId", Value: userId}})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if err = cursor.All(ctx, &results); err != nil {
		log.Println(err)
		return nil, err
	}

	return results, nil
}

func (repo *MemoRepoMongoDB) GetByDateRange(start, end, userId string) ([]entity.Memo, error) {
	ctx, cancel, collection := getCollection(repo, repo.dbMame, repo.collection)
	defer cancel()

	//"UserId = :userId and MonthDay BETWEEN :start and :end"
	filter := bson.D{{Key: "userId", Value: userId},
		{Key: "monthDay", Value: bson.D{{Key: "$gte", Value: start}}},
		{Key: "monthDay", Value: bson.D{{Key: "$lte", Value: end}}},
	}
	var results []entity.Memo
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if err = cursor.All(ctx, &results); err != nil {
		log.Println(err)
		return nil, err
	}

	return results, nil
}

func (repo *MemoRepoMongoDB) Update(memo entity.Memo) error {
	ctx, cancel, collection := getCollection(repo, repo.dbMame, repo.collection)
	defer cancel()

	filter := bson.D{{Key: "userId", Value: memo.UserId}, {Key: "_id", Value: memo.Id}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "subject", Value: memo.Subject}}},
		{Key: "$set", Value: bson.D{{Key: "description", Value: memo.Description}}},
		{Key: "$set", Value: bson.D{{Key: "monthDay", Value: strconv.Itoa(memo.MonthDay)}}},
		{Key: "$set", Value: bson.D{{Key: "startYear", Value: strconv.Itoa(memo.StartYear)}}},
		{Key: "$set", Value: bson.D{{Key: "lunar", Value: memo.Lunar}}},
		{Key: "$set", Value: bson.D{{Key: "lastModifiedTime", Value: strconv.FormatInt(time.Now().UTC().Unix(), 10)}}},
	}

	err := collection.FindOneAndUpdate(ctx, filter, update).Err()

	if err != nil {
		log.Printf("Got error calling UpdateItem: %s", err)
		return err
	}

	return nil
}

func (repo *MemoRepoMongoDB) Delete(id string) error {
	ctx, cancel, collection := getCollection(repo, repo.dbMame, repo.collection)
	defer cancel()

	filter := bson.D{{Key: "_id", Value: id}}
	err := collection.FindOneAndDelete(ctx, filter).Err()
	if err != nil {
		log.Printf("Got error calling UpdateItem: %s", err)
		return err
	}

	return nil
}

func getCollection(repo *MemoRepoMongoDB, dbName, collectionName string) (context.Context, context.CancelFunc, *mongo.Collection) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := mongo.Connect(ctx, repo.ClientOptions)
	if err != nil {
		log.Println(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Println(err)
	}

	collection := client.Database(dbName).Collection(collectionName)
	return ctx, cancel, collection
}
