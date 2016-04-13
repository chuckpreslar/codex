package codex

import (
	"github.com/chuckpreslar/codex/lib/table"
)

// TableFor ...
func TableFor(name string) table.Table {
	return table.InitializeTable(name)
}
