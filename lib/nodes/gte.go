package nodes

import (
	"github.com/chuckpreslar/codex/lib/aux"
	"github.com/chuckpreslar/codex/lib/interfaces"
)

// Gte ...
type Gte struct {
	Left  interfaces.Acceptor
	Right interfaces.Acceptor
}

// Or ...
func (g *Gte) Or(any interface{}) *Or {
	return InitializeOr(g, aux.CreateAcceptorFrom(any))
}

// And ...
func (g *Gte) And(any interface{}) *And {
	return InitializeAnd(g, aux.CreateAcceptorFrom(any))
}

// Accept ...
func (g *Gte) Accept(visitor interfaces.Visitor) string {
	return visitor.Visit(g)
}

// InitializeGte ...
func InitializeGte(left, right interfaces.Acceptor) *Gte {
	var gte = new(Gte)

	gte.Left = left
	gte.Right = right

	return gte
}
