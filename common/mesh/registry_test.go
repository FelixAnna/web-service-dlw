package mesh

import (
	"os"
	"testing"

	"github.com/FelixAnna/web-service-dlw/common/aws"
	"github.com/FelixAnna/web-service-dlw/common/mocks"
	"github.com/stretchr/testify/assert"
)

var service *Registry
var helper aws.AwsInterface

func init() {
	helper = &mocks.MockAwsHelper{}
}

func TestProvideRegistry(t *testing.T) {
	service = ProvideRegistry(aws.ProvideAWSService(helper))

	assert.NotNil(t, service)
	assert.Empty(t, service.consulRegAddr)
}

func TestProvideRegistryProd(t *testing.T) {
	os.Setenv("profile", "prod")

	service = ProvideRegistry(aws.ProvideAWSService(helper))

	assert.NotNil(t, service)
	assert.NotEmpty(t, service.consulRegAddr)
}

func TestProvideRegistryDev(t *testing.T) {
	os.Setenv("profile", "dev")

	service = ProvideRegistry(aws.ProvideAWSService(helper))

	assert.NotNil(t, service)
	assert.NotEmpty(t, service.consulRegAddr)
}

func TestGetRegistryConsul(t *testing.T) {
	os.Setenv("profile", "prod")

	service = ProvideRegistry(aws.ProvideAWSService(helper))
	result := service.GetRegistry()
	assert.NotNil(t, result)
}

func SkipTestGetRegistryOther(t *testing.T) {
	os.Setenv("profile", "Local")

	service = ProvideRegistry(aws.ProvideAWSService(helper))
	result := service.GetRegistry()

	//need mock kubernetes.NewRegistry
	assert.NotNil(t, result)
}
