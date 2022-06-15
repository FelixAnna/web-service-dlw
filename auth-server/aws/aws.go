package aws

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

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

var sess = session.Must(session.NewSessionWithOptions(session.Options{
	SharedConfigState: session.SharedConfigEnable,
}))

var Parameters map[string]string

func init() {
	ssmClient := ssm.New(sess)

	path := "/dlf/dev"
	out, err := ssmClient.GetParametersByPath(&ssm.GetParametersByPathInput{
		Path:           &path,
		WithDecryption: aws.Bool(true),
		Recursive:      aws.Bool(true),
	})

	if err != nil {
		log.Printf("Error when geting ssm parameters: %v", err)
		panic(err)
	}

	Parameters = make(map[string]string, len(out.Parameters))
	for _, parameter := range out.Parameters {
		Parameters[*parameter.Name] = *parameter.Value
	}
}
