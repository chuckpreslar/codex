package nodes

import (
	"github.com/chuckpreslar/codex/lib/aux"
	"github.com/chuckpreslar/codex/lib/interfaces"
)

// Neq ...
type Neq struct {
	Left  interfaces.Acceptor
	Right interfaces.Acceptor
}

// Or ...
func (n *Neq) Or(any interface{}) *Or {
	return InitializeOr(n, aux.CreateAcceptorFrom(any))
}

// And ...
func (n *Neq) And(any interface{}) *And {
	return InitializeAnd(n, aux.CreateAcceptorFrom(any))
}

// Accept ...
func (n *Neq) Accept(visitor interfaces.Visitor) string {
	return visitor.Visit(n)
}

// InitializeNeq ...
func InitializeNeq(left, right interfaces.Acceptor) *Neq {
	var neq = new(Neq)

	neq.Left = left
	neq.Right = right

	return neq
}
