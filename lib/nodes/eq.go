package nodes

import (
	"github.com/chuckpreslar/codex/lib/aux"
	"github.com/chuckpreslar/codex/lib/interfaces"
)

// Eq ...
type Eq struct {
	Left  interfaces.Acceptor
	Right interfaces.Acceptor
}

// Or ...
func (e *Eq) Or(any interface{}) *Or {
	return InitializeOr(e, aux.CreateAcceptorFrom(any))
}

// And ...
func (e *Eq) And(any interface{}) *And {
	return InitializeAnd(e, aux.CreateAcceptorFrom(any))
}

// Accept ...
func (e *Eq) Accept(visitor interfaces.Visitor) string {
	return visitor.Visit(e)
}

// InitializeEq ...
func InitializeEq(left, right interfaces.Acceptor) *Eq {
	var eq = new(Eq)

	eq.Left = left
	eq.Right = right

	return eq
}
