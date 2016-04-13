package cores

import "github.com/chuckpreslar/codex/lib/interfaces"

// Select ...
type Select struct {
	Source struct {
		Left  interfaces.Acceptor
		Right []interfaces.Acceptor
	}
	Projections []interfaces.Acceptor
	Wheres      interfaces.Acceptor
	Groups      interfaces.Acceptor
	Havings     interfaces.Acceptor
	Windows     interfaces.Acceptor
}

// InitializeSelect ...
func InitializeSelect() *Select {
	var sel = new(Select)
	return sel
}
