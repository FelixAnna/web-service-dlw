package snowflake

import (
	"fmt"
	"os"
	"strconv"

	"github.com/bwmarrin/snowflake"
)

var node *snowflake.Node

// InitSnowflake initiate Snowflake node singleton.
func InitSnowflake() error {
	// Get node number from env DLW_NODE_NO
	key, ok := os.LookupEnv("DLW_NODE_NO")
	if !ok {
		return fmt.Errorf("DLW_NODE_NO is not set in system environment")
	}

	// Parse node number
	nodeNo, err := strconv.ParseInt(key, 10, 64)
	if err != nil {
		return err
	}

	// Create snowflake node
	n, err := snowflake.NewNode(nodeNo)
	if err != nil {
		return err
	}

	// Set node
	node = n
	return nil
}

// GenerateSnowflake generate Twitter Snowflake ID
func GenerateSnowflake() string {
	return node.Generate().String()
}
