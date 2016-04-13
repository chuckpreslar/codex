package codex

import (
	"github.com/chuckpreslar/codex/lib"
)

// TableFor ...
func TableFor(name string) table.Table {
	return table.InitializeTable(name)
}
