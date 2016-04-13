package nodes

import (
	"github.com/chuckpreslar/codex/lib/aux"
	"github.com/chuckpreslar/codex/lib/interfaces"
)

// And ...
type And struct {
	Left  interfaces.Acceptor
	Right interfaces.Acceptor
}

// Or ...
func (a *And) Or(any interface{}) *Or {
	return InitializeOr(a, aux.CreateAcceptorFrom(any))
}

// And ...
func (a *And) And(any interface{}) *And {
	return InitializeAnd(a, aux.CreateAcceptorFrom(any))
}

// Accept ...
func (a *And) Accept(visitor interfaces.Visitor) string {
	return visitor.Visit(a)
}

// InitializeAnd ...
func InitializeAnd(left, right interfaces.Acceptor) *And {
	var and = new(And)

	and.Left = left
	and.Right = right

	return and
}
