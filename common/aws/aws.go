package aws

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

const basePath = "/dlf"

type AWSService struct {
	sess       *session.Session
	parameters map[string]string
}

func ProvideAWSService(helper AwsInterface) *AWSService {
	sess := helper.CreateSess()
	parameters := helper.LoadParameters(sess)
	return &AWSService{sess: sess, parameters: parameters}
}

func (service *AWSService) GetParameterByKey(key string) string {
	env := "dev"
	fullKey := fmt.Sprintf("%v/%v/%v", basePath, env, key)
	return service.parameters[fullKey]
}

func (service *AWSService) GetDynamoDBClient() *dynamodb.DynamoDB {
	dynamoDB := dynamodb.New(service.sess)
	return dynamoDB
}
