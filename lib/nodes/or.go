package nodes

import (
	"github.com/chuckpreslar/codex/lib/aux"
	"github.com/chuckpreslar/codex/lib/interfaces"
)

// Or ...
type Or struct {
	Left  interfaces.Acceptor
	Right interfaces.Acceptor
}

// Or ...
func (o *Or) Or(any interface{}) *Or {
	return InitializeOr(o, aux.CreateAcceptorFrom(any))
}

// And ...
func (o *Or) And(any interface{}) *And {
	return InitializeAnd(o, aux.CreateAcceptorFrom(any))
}

// Accept ...
func (o *Or) Accept(visitor interfaces.Visitor) string {
	return visitor.Visit(o)
}

// InitializeOr ...
func InitializeOr(left, right interfaces.Acceptor) *Or {
	var or = new(Or)

	or.Left = left
	or.Right = right

	return or
}
