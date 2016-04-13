package base

import (
	"github.com/chuckpreslar/codex/lib/interfaces"
)

// Nil ...
type Nil struct{}

// Accept ..
func (n *Nil) Accept(visitor interfaces.Visitor) string {
	return visitor.Visit(n)
}

// InitializeNil ...
func InitializeNil() *Nil {
	return new(Nil)
}
