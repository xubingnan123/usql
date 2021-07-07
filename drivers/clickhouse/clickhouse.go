package clickhouse

import (
	// DRIVER: clickhouse
	_ "github.com/kshvakov/clickhouse"

	"github.com/xubingnan123/usql/drivers"
)

func init() {
	drivers.Register("clickhouse", drivers.Driver{
		AllowMultilineComments: true,
	})
}
