package nodes

import (
	"github.com/chuckpreslar/codex/lib/aux"
	"github.com/chuckpreslar/codex/lib/interfaces"
)

// Lt ...
type Lt struct {
	Left  interfaces.Acceptor
	Right interfaces.Acceptor
}

// Or ...
func (l *Lt) Or(any interface{}) *Or {
	return InitializeOr(l, aux.CreateAcceptorFrom(any))
}

// And ...
func (l *Lt) And(any interface{}) *And {
	return InitializeAnd(l, aux.CreateAcceptorFrom(any))
}

// Accept ...
func (l *Lt) Accept(visitor interfaces.Visitor) string {
	return visitor.Visit(l)
}

// InitializeLt ...
func InitializeLt(left, right interfaces.Acceptor) *Lt {
	var lt = new(Lt)

	lt.Right = right
	lt.Left = left

	return lt
}
