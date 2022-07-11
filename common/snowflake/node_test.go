package snowflake

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitSnowflakeNoEnv(t *testing.T) {
	//os.Setenv("DLW_NODE_NO", "")
	err := InitSnowflake()
	assert.NotNil(t, err)
	assert.EqualValues(t, err.Error(), "DLW_NODE_NO is not set in system environment")
}

func TestInitSnowflakeInvalidEnv(t *testing.T) {
	os.Setenv("DLW_NODE_NO", "NaN")
	err := InitSnowflake()
	assert.NotNil(t, err)
	assert.NotEqualValues(t, err.Error(), "DLW_NODE_NO is not set in system environment")
}

func TestInitSnowflakeOutOfRangeEnv(t *testing.T) {
	os.Setenv("DLW_NODE_NO", "1025")
	err := InitSnowflake()
	assert.NotNil(t, err)
	assert.NotEqualValues(t, err.Error(), "DLW_NODE_NO is not set in system environment")
}

func TestInitSnowflakeOk(t *testing.T) {
	os.Setenv("DLW_NODE_NO", "1023")
	err := InitSnowflake()
	assert.Nil(t, err)
	assert.NotNil(t, node)
}

func TestGenerateSnowflake(t *testing.T) {
	os.Setenv("DLW_NODE_NO", "1023")
	InitSnowflake()

	result := GenerateSnowflake()

	assert.NotEmpty(t, result)
}
