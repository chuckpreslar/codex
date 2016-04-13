package base

import (
	"github.com/chuckpreslar/codex/lib/interfaces"
)

// Boolean ...
type Boolean struct {
	Value bool
}

// Accept ..
func (b *Boolean) Accept(visitor interfaces.Visitor) string {
	return visitor.Visit(b)
}

// InitializeBoolean ...
func InitializeBoolean(value bool) *Boolean {
	var boolean = new(Boolean)

	boolean.Value = value

	return boolean
}
