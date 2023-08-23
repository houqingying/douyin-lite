package snowflake

import (
	"sync"
	"time"

	"github.com/bwmarrin/snowflake"
)

var (
	once sync.Once
	node *snowflake.Node
)

func InitSnowflakeNode(startTime string, machineID int64) error {
	var err error

	once.Do(func() {
		var st time.Time
		st, err = time.Parse("2006-01-02", startTime)
		if err != nil {
			return
		}
		snowflake.Epoch = st.UnixNano() / 1000000
		node, err = snowflake.NewNode(machineID)
	})

	return err
}

func GenerateID() int64 {
	return node.Generate().Int64()
}
