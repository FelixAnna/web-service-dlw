package aws

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/ssm"
)

/*
func GetClientByConfig() *dynamodb.Client {
	// Using the SDK's default configuration, loading additional config
	// and credentials values from the environment variables, shared
	// credentials, and shared configuration files
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("ap-southeast-1"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	// Using the Config value, create the DynamoDB client
	client := dynamodb.NewFromConfig(cfg)

	return client
}*/

const basePath = "/dlf"

var (
	sess       *session.Session
	parameters map[string]string
)

func init() {
	sess = session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	ssmClient := ssm.New(sess)

	out, err := ssmClient.GetParametersByPath(&ssm.GetParametersByPathInput{
		Path:           aws.String(basePath),
		WithDecryption: aws.Bool(true),
		Recursive:      aws.Bool(true),
	})

	if err != nil {
		log.Fatalf("Error when geting ssm parameters: %v", err)
		panic(err)
	}

	parameters = make(map[string]string, len(out.Parameters))
	for _, parameter := range out.Parameters {
		parameters[*parameter.Name] = *parameter.Value
	}
}

func GetParameterByKey(key string) string {
	env := "dev"
	fullKey := fmt.Sprintf("%v/%v/%v", basePath, env, key)
	return parameters[fullKey]
}

func GetDynamoDBClient() *dynamodb.DynamoDB {
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials
	// and region from the shared configuration file ~/.aws/config.

	//set AWS_REGION=ap-southeast-1
	/*sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-southeast-1")},
	)

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("ap-southeast-1"),
		Credentials: credentials.NewStaticCredentials(key, value, ""),
	})

	if err != nil {
		log.Fatalf("Error when connecting to dynamodb: %v", err)
	}*/

	// Create DynamoDB client
	dynamoDB := dynamodb.New(sess)

	return dynamoDB
}
