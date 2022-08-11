package aws

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var basePath string = "/dlf"

func init() {
	env := os.Getenv("profile")
	if env != "prod" {
		env = "dev"
	}

	basePath = fmt.Sprintf("%v/%v", basePath, env)
}

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
	fullKey := fmt.Sprintf("%v/%v", basePath, key)
	return service.parameters[fullKey]
}

func (service *AWSService) GetDynamoDBClient() *dynamodb.DynamoDB {
	dynamoDB := dynamodb.New(service.sess)
	return dynamoDB
}
