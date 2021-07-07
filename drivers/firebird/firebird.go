package firebird

import (
	// DRIVER: firebirdsql
	_ "github.com/nakagami/firebirdsql"

	"github.com/xubingnan123/usql/drivers"
)

func init() {
	drivers.Register("firebirdsql", drivers.Driver{
		AllowMultilineComments: true,
	})
}
