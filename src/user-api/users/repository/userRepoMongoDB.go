package repository

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/FelixAnna/web-service-dlw/common/aws"
	"github.com/FelixAnna/web-service-dlw/user-api/users/entity"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var MongoRepoSet = wire.NewSet(ProvideUserRepoMongoDB, wire.Bind(new(UserRepo), new(*UserRepoMongoDB)))

type UserRepoMongoDB struct {
	ClientOptions *options.ClientOptions
	collection    string
	dbMame        string
}

func ProvideUserRepoMongoDB(awsService *aws.AWSService) *UserRepoMongoDB {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)

	clientOptions := options.Client().
		ApplyURI(awsService.GetParameterByKey("mongo/connectionstring")).
		SetServerAPIOptions(serverAPIOptions)

	return &UserRepoMongoDB{
		ClientOptions: clientOptions,
		dbMame:        "dlw_memo",
		collection:    "users",
	}
}

func (repo *UserRepoMongoDB) GetAllTables() {
	log.Println(repo.collection)
}

func (repo *UserRepoMongoDB) GetAll() ([]entity.User, error) {
	ctx, cancel, collection := getCollection(repo, repo.dbMame, repo.collection)
	defer cancel()

	var results []entity.User
	cursor, err := collection.Find(ctx, bson.D{})
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

func (repo *UserRepoMongoDB) Add(user *entity.User) (*string, error) {
	if eu, _ := repo.GetByEmail(user.Email); eu != nil {
		return nil, errors.New("user with same email already exists")
	}

	randId := fmt.Sprintf("%d%03d", time.Now().Unix(), rand.Intn(1000))
	user.Id = randId
	user.CreateTime = strconv.FormatInt(time.Now().UTC().Unix(), 10)

	ctx, cancel, collection := getCollection(repo, repo.dbMame, repo.collection)
	defer cancel()

	_, err := collection.InsertOne(ctx, *user)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &user.Id, nil
}

func (repo *UserRepoMongoDB) GetById(userId string) (*entity.User, error) {
	ctx, cancel, collection := getCollection(repo, repo.dbMame, repo.collection)
	defer cancel()

	var result entity.User //= entity.Questions{}
	err := collection.FindOne(ctx, bson.D{{Key: "_id", Value: userId}}).Decode(&result)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &result, nil
}

func (repo *UserRepoMongoDB) GetByEmail(email string) (*entity.User, error) {
	ctx, cancel, collection := getCollection(repo, repo.dbMame, repo.collection)
	defer cancel()

	var result entity.User
	err := collection.FindOne(ctx, bson.D{{Key: "email", Value: email}}).Decode(&result)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &result, nil
}

func (repo *UserRepoMongoDB) UpdateBirthday(userId, birthday string) error {
	ctx, cancel, collection := getCollection(repo, repo.dbMame, repo.collection)
	defer cancel()

	filter := bson.D{{Key: "_id", Value: userId}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "birthday", Value: birthday}}}}

	err := collection.FindOneAndUpdate(ctx, filter, update).Err()

	if err != nil {
		log.Printf("Got error calling UpdateItem: %s", err)
		return err
	}

	return nil
}

func (repo *UserRepoMongoDB) UpdateAddress(userId string, addresses []entity.Address) error {
	ctx, cancel, collection := getCollection(repo, repo.dbMame, repo.collection)
	defer cancel()

	filter := bson.D{{Key: "_id", Value: userId}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "address", Value: addresses}}}}

	err := collection.FindOneAndUpdate(ctx, filter, update).Err()

	if err != nil {
		log.Printf("Got error calling UpdateItem: %s", err)
		return err
	}

	return nil
}

func (repo *UserRepoMongoDB) Delete(userId string) error {
	ctx, cancel, collection := getCollection(repo, repo.dbMame, repo.collection)
	defer cancel()

	filter := bson.D{{Key: "_id", Value: userId}}
	err := collection.FindOneAndDelete(ctx, filter).Err()
	if err != nil {
		log.Printf("Got error calling UpdateItem: %s", err)
		return err
	}

	return nil
}

func getCollection(repo *UserRepoMongoDB, dbName, collectionName string) (context.Context, context.CancelFunc, *mongo.Collection) {
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
