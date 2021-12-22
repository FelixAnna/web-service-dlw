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

	config "github.com/FelixAnna/web-service-dlw/common/aws"
	"github.com/FelixAnna/web-service-dlw/memo-api/memo/entity"
)

var (
	tableName string
	client    *dynamodb.DynamoDB
)

func init() {
	tableName = "dlf.Memos"
	client = config.GetDynamoDBClient()
}

type MemoRepoDynamoDB struct {
	MemoRepo
}

func (m *MemoRepoDynamoDB) Add(memo *entity.Memo) (*string, error) {
	randId := fmt.Sprintf("%d%03d", time.Now().Unix(), rand.Intn(1000))
	memo.Id = randId
	memo.CreateTime = strconv.FormatInt(time.Now().UTC().Unix(), 10)

	memoJson, err := dynamodbattribute.MarshalMap(memo)
	if err != nil {
		log.Printf("Got error marshalling new memo item: %s", err)
		return nil, err
	}

	_, err = client.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      memoJson,
	})

	if err != nil {
		log.Printf("Got error calling PutItem: %s", err)
		return nil, err
	}

	return &memo.Id, nil
}

func (m *MemoRepoDynamoDB) GetById(id string) (*entity.Memo, error) {
	result, err := client.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"Id": {S: aws.String(id)},
		},
	},
	)

	if err != nil {
		log.Printf("Got error calling GetItem: %s", err)
		return nil, err
	}

	if result.Item == nil {
		msg := "Could not find memo with Id: '" + id + "'"
		return nil, errors.New(msg)
	}

	item := entity.Memo{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		log.Printf("Failed to unmarshal Record, %v", err)
		return nil, err
	}

	return &item, nil
}

func (m *MemoRepoDynamoDB) GetByUserId(userId string) ([]entity.Memo, error) {
	result, err := client.Query(&dynamodb.QueryInput{
		TableName:              aws.String(tableName),
		IndexName:              aws.String("UserId-MonthDay-index"),
		KeyConditionExpression: aws.String("UserId = :userId"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":userId": {S: &userId},
		},
		Limit: aws.Int64(100),
	})

	if err != nil {
		log.Printf("Query API call failed: %s", err)
		return nil, err
	}

	var memos []entity.Memo = make([]entity.Memo, len(result.Items))
	for index, item := range result.Items {
		memo := entity.Memo{}
		err = dynamodbattribute.UnmarshalMap(item, &memo)

		if err != nil {
			log.Printf("Got error unmarshalling to memo: %s", err)
			return nil, err
		}

		memos[index] = memo
	}

	return memos, nil
}

func (m *MemoRepoDynamoDB) GetByDateRange(start, end, userId string) ([]entity.Memo, error) {
	result, err := client.Query(&dynamodb.QueryInput{
		TableName:              aws.String(tableName),
		IndexName:              aws.String("UserId-MonthDay-index"),
		KeyConditionExpression: aws.String("UserId = :userId and MonthDay BETWEEN :start and :end"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":start":  {N: &start},
			":end":    {N: &end},
			":userId": {S: &userId},
		},
		Limit: aws.Int64(100),
	})

	if err != nil {
		log.Printf("Query API call failed: %s", err)
		return nil, err
	}

	var memos []entity.Memo = make([]entity.Memo, len(result.Items))
	for index, item := range result.Items {
		memo := entity.Memo{}
		err = dynamodbattribute.UnmarshalMap(item, &memo)

		if err != nil {
			log.Printf("Got error unmarshalling to memo: %s", err)
			return nil, err
		}

		memos[index] = memo
	}

	return memos, nil
}

func (m *MemoRepoDynamoDB) Update(memo entity.Memo) error {
	oldMemo, err := m.GetById(memo.Id)
	if err != nil {
		return errors.New("memo not exists")
	}

	if oldMemo.UserId != memo.UserId {
		return fmt.Errorf("you are not owner of the memo id: %v, owner: %v, request userId: %v", memo.Id, oldMemo.UserId, memo.UserId)
	}

	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":subject":    {S: aws.String(memo.Subject)},
			":desc":       {S: aws.String(memo.Description)},
			":monthDay":   {N: aws.String(strconv.Itoa(memo.MonthDay))},
			":year":       {N: aws.String(strconv.Itoa(memo.StartYear))},
			":lunar":      {BOOL: aws.Bool(memo.Lunar)},
			":updateTime": {S: aws.String(strconv.FormatInt(time.Now().UTC().Unix(), 10))},
		},
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"Id": {S: aws.String(memo.Id)},
		},
		ReturnValues:     aws.String("ALL_NEW"),
		UpdateExpression: aws.String("set Subject = :subject, Description = :desc, MonthDay = :monthDay, StartYear = :year, Lunar = :lunar, LastModifiedTime = :updateTime"),
	}

	_, err = client.UpdateItem(input)
	if err != nil {
		log.Printf("Got error calling UpdateItem: %s", err)
		return err
	}

	return nil
}

func (m *MemoRepoDynamoDB) Delete(id string) error {
	if _, err := m.GetById(id); err != nil {
		return errors.New("memo not exists")
	}

	_, err := client.DeleteItem(&dynamodb.DeleteItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"Id": {S: aws.String(id)},
		},
	})

	if err != nil {
		log.Printf("Got error calling UpdateItem: %s", err)
		return err
	}

	return nil
}
