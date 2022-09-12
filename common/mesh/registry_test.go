package mesh

import (
	"os"
	"testing"

	"github.com/FelixAnna/web-service-dlw/common/aws"
	"github.com/FelixAnna/web-service-dlw/common/mocks"
	"github.com/stretchr/testify/assert"
)

var service *Registry

func init() {
	helper := mocks.MockAwsHelper{}
	service = ProvideRegistry(aws.ProvideAWSService(&helper))
}

func TestProvideRegistry(t *testing.T) {
	assert.NotNil(t, service)
	assert.NotEmpty(t, service.consulRegAddr)
}

func TestGetRegistryProd(t *testing.T) {
	os.Setenv("profile", "prod")

	result := service.GetRegistry()

	assert.NotNil(t, result)
}

func TestGetRegistryDev(t *testing.T) {
	os.Setenv("profile", "dev")

	result := service.GetRegistry()

	assert.NotNil(t, result)
}

func SkipTestGetRegistryOther(t *testing.T) {
	os.Setenv("profile", "Local")

	result := service.GetRegistry()

	//need mock kubernetes.NewRegistry
	assert.NotNil(t, result)
}
