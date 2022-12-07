package aws

import (
	"log"

	"github.com/FelixAnna/web-service-dlw/common/mocks"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/google/wire"
)

var AwsMockSet = wire.NewSet(mocks.ProvideMockAwsHelper, wire.Bind(new(AwsInterface), new(*mocks.MockAwsHelper)))
var AwsSet = wire.NewSet(ProvideAwsHelper, wire.Bind(new(AwsInterface), new(*AwsHelper)))

type AwsInterface interface {
	CreateSess() *session.Session
	LoadParameters(sess *session.Session) map[string]string
}

type AwsHelper struct {
}

func ProvideAwsHelper() *AwsHelper {
	return &AwsHelper{}
}

/*
func GetClientByConfig() *dynamodb.Client {
	// Using the SDK's default configuration, loading additional config
	// and credentials values from the environment variables, shared
	// credentials, and shared configuration files
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("ap-southeast-1"))
	if err != nil {
		log.Printf("unable to load SDK config, %v", err)
	}

	// Using the Config value, create the DynamoDB client
	client := dynamodb.NewFromConfig(cfg)

	return client
}*/

func (helper *AwsHelper) CreateSess() *session.Session {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	return sess
}

func (helper *AwsHelper) LoadParameters(sess *session.Session) map[string]string {
	parameters := make(map[string]string)

	ssmClient := ssm.New(sess)
	input := ssm.GetParametersByPathInput{
		MaxResults:     aws.Int64(10),
		Path:           aws.String(basePath),
		WithDecryption: aws.Bool(true),
		Recursive:      aws.Bool(true),
	}

	if err := ssmClient.GetParametersByPathPages(&input, func(gpbpo *ssm.GetParametersByPathOutput, b bool) bool {
		for _, parameter := range gpbpo.Parameters {
			parameters[*parameter.Name] = *parameter.Value
		}
		return true
	}); err != nil {
		log.Printf("Error when geting ssm parameters: %v", err)
		panic(err)
	}

	return parameters
}
