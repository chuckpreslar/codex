package nodes

import (
	"github.com/chuckpreslar/codex/lib/aux"
	"github.com/chuckpreslar/codex/lib/interfaces"
)

// Lte ...
type Lte struct {
	Left  interfaces.Acceptor
	Right interfaces.Acceptor
}

// Or ...
func (l *Lte) Or(any interface{}) *Or {
	return InitializeOr(l, aux.CreateAcceptorFrom(any))
}

// And ...
func (l *Lte) And(any interface{}) *And {
	return InitializeAnd(l, aux.CreateAcceptorFrom(any))
}

// Accept ...
func (l *Lte) Accept(visitor interfaces.Visitor) string {
	return visitor.Visit(l)
}

// InitializeLte ...
func InitializeLte(left, right interfaces.Acceptor) *Lte {
	var lte = new(Lte)

	lte.Left = left
	lte.Right = right

	return lte
}
