package snowf

import (
	"github.com/bwmarrin/snowflake"
	"github.com/google/wire"
	appConfig "github.com/oa-meeting/pkg/config"
	"time"
)

var node *snowflake.Node
var Provider = wire.NewSet(NewSf)

func NewSf() *snowflake.Node {
	var err error
	var st time.Time
	st, err = time.Parse("2006-01-02", appConfig.Data.SnowFlake.StartTime)
	if err != nil {
		panic(err)
	}
	snowflake.Epoch = st.UnixNano() / 1000000
	node, err = snowflake.NewNode(int64(appConfig.Data.SnowFlake.NodeNum))
	return node
}
func GenID() int64 {
	return node.Generate().Int64()
}
