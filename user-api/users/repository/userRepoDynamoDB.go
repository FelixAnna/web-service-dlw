package repository

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/google/wire"

	config "github.com/FelixAnna/web-service-dlw/common/aws"
	"github.com/FelixAnna/web-service-dlw/user-api/users/entity"
)

var RepoSet = wire.NewSet(ProvideUserRepoDynamoDB, wire.Bind(new(UserRepo), new(*UserRepoDynamoDB)))

type UserRepoDynamoDB struct {
	TableName string
	Client    *dynamodb.DynamoDB
}

func ProvideUserRepoDynamoDB(awsService *config.AWSService) *UserRepoDynamoDB {
	return &UserRepoDynamoDB{Client: awsService.GetDynamoDBClient(), TableName: "dlf.Users"}
}

func (u *UserRepoDynamoDB) GetAllTables() {
	// Build the request with its input parameters
	resp, err := u.Client.ListTables(&dynamodb.ListTablesInput{
		Limit: aws.Int64(5),
	})
	if err != nil {
		log.Printf("failed to list tables, %v", err)
	}

	log.Println("Tables:")

	for _, tableName := range resp.TableNames {
		log.Println(tableName)
	}
}

func (u *UserRepoDynamoDB) GetAll() ([]entity.User, error) {
	filt := expression.Name("Email").AttributeExists()
	projection := expression.NamesList(expression.Name("Id"), expression.Name("Name"), expression.Name("Email"), expression.Name("Phone"), expression.Name("Birthday"))

	expr, err := expression.NewBuilder().WithFilter(filt).WithProjection(projection).Build()
	if err != nil {
		log.Printf("Got error building expression: %s", err)
		return nil, err
	}

	// Build the query input parameters
	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String(u.TableName),
	}

	result, err := u.Client.Scan(params)
	if err != nil {
		log.Printf("Query API call failed: %s", err)
		return nil, err
	}

	var users []entity.User = make([]entity.User, 0)
	for _, item := range result.Items {
		user := entity.User{}

		err = dynamodbattribute.UnmarshalMap(item, &user)

		if err != nil {
			log.Printf("Got error unmarshalling: %s", err)
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (u *UserRepoDynamoDB) Add(user *entity.User) (*string, error) {
	if eu, _ := u.GetByEmail(user.Email); eu != nil {
		return nil, errors.New("user with same email already exists")
	}

	randId := fmt.Sprintf("%d%03d", time.Now().Unix(), rand.Intn(1000))
	user.Id = randId
	user.CreateTime = strconv.FormatInt(time.Now().UTC().Unix(), 10)

	userJson, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		log.Printf("Got error marshalling new User item: %s", err)
		return nil, err
	}

	_, err = u.Client.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(u.TableName),
		Item:      userJson,
	})

	if err != nil {
		log.Printf("Got error calling PutItem: %s", err)
		return nil, err
	}

	return &user.Id, nil
}

func (u *UserRepoDynamoDB) GetById(userId string) (*entity.User, error) {
	result, err := u.Client.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(u.TableName),
		Key: map[string]*dynamodb.AttributeValue{
			"Id": {S: aws.String(userId)},
		},
	},
	)

	if err != nil {
		log.Printf("Got error calling GetItem: %s", err)
		return nil, err
	}

	if result.Item == nil {
		msg := "Could not find user with Id: '" + userId + "'"
		return nil, errors.New(msg)
	}

	item := entity.User{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		log.Printf("Failed to unmarshal Record, %v", err)
		return nil, err
	}

	return &item, nil
}

func (u *UserRepoDynamoDB) GetByEmail(email string) (*entity.User, error) {
	result, err := u.Client.Query(&dynamodb.QueryInput{
		TableName:              aws.String(u.TableName),
		IndexName:              aws.String("Email-index"),
		KeyConditionExpression: aws.String("Email = :email"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":email": {S: &email},
		},
		Limit: aws.Int64(1),
	})

	if err != nil {
		log.Printf("Query API call failed: %s", err)
		return nil, err
	}

	if length := len(result.Items); length > 0 {
		user := entity.User{}

		err = dynamodbattribute.UnmarshalMap(result.Items[0], &user)

		if err != nil {
			log.Printf("Got error unmarshalling: %s", err)
			return nil, err
		}

		return &user, nil

	} else {
		return nil, err
	}
}

func (u *UserRepoDynamoDB) UpdateBirthday(userId, birthday string) error {
	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":birthday": {S: aws.String(birthday)},
		},
		TableName: aws.String(u.TableName),
		Key: map[string]*dynamodb.AttributeValue{
			"Id": {S: aws.String(userId)},
		},
		ReturnValues:     aws.String("ALL_NEW"),
		UpdateExpression: aws.String("set Birthday = :birthday"),
	}

	_, err := u.Client.UpdateItem(input)
	if err != nil {
		log.Printf("Got error calling UpdateItem: %s", err)
		return err
	}

	return nil
}

func (u *UserRepoDynamoDB) UpdateAddress(userId string, addresses []entity.Address) error {

	addressJson, err := dynamodbattribute.MarshalList(addresses)
	if err != nil {
		log.Printf("Got error calling UpdateItem: %s", err)
		return err
	}

	log.Println(addressJson)
	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":addresses": {L: addressJson},
		},
		TableName: aws.String(u.TableName),
		Key: map[string]*dynamodb.AttributeValue{
			"Id": {S: aws.String(userId)},
		},
		ReturnValues:     aws.String("ALL_NEW"),
		UpdateExpression: aws.String("set Address = :addresses"),
	}

	_, err = u.Client.UpdateItem(input)
	if err != nil {
		log.Printf("Got error calling UpdateItem: %s", err)
		return err
	}

	return nil
}

func (u *UserRepoDynamoDB) Delete(userId string) error {
	if _, err := u.GetById(userId); err != nil {
		return errors.New("user not exists")
	}

	_, err := u.Client.DeleteItem(&dynamodb.DeleteItemInput{
		TableName: aws.String(u.TableName),
		Key: map[string]*dynamodb.AttributeValue{
			"Id": {S: aws.String(userId)},
		},
	})

	if err != nil {
		log.Printf("Got error calling UpdateItem: %s", err)
		return err
	}

	return nil
}
