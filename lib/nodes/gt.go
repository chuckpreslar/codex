package nodes

import (
	"github.com/chuckpreslar/codex/lib/aux"
	"github.com/chuckpreslar/codex/lib/interfaces"
)

// Gt ...
type Gt struct {
	Left  interfaces.Acceptor
	Right interfaces.Acceptor
}

// Or ...
func (g *Gt) Or(any interface{}) *Or {
	return InitializeOr(g, aux.CreateAcceptorFrom(any))
}

// And ...
func (g *Gt) And(any interface{}) *And {
	return InitializeAnd(g, aux.CreateAcceptorFrom(any))
}

// Accept ...
func (g *Gt) Accept(visitor interfaces.Visitor) string {
	return visitor.Visit(g)
}

// InitializeGt ...
func InitializeGt(left, right interfaces.Acceptor) *Gt {
	var gt = new(Gt)

	gt.Left = left
	gt.Right = right

	return gt
}
