package aws

import (
	"testing"

	test "github.com/FelixAnna/web-service-dlw/common/testing"
	"github.com/stretchr/testify/assert"
)

var service *AWSService

func init() {
	helper := test.MockAwsHelper{}
	service = ProvideAWSService(&helper)
}

func TestProvideAWSService(t *testing.T) {
	assert.NotNil(t, service)
	assert.NotNil(t, service.sess)
	assert.Greater(t, len(service.parameters), 0)
}

func TestGetParameterByKey(t *testing.T) {
	key := "key1"

	result := service.GetParameterByKey(key)

	assert.NotNil(t, result)
	assert.Equal(t, result, "value1")
}

func TestGetDynamoDBClient(t *testing.T) {
	result := service.GetDynamoDBClient()

	assert.NotNil(t, result)
}
