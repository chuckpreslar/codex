package base

import (
	"github.com/chuckpreslar/codex/lib/interfaces"
)

// Integer ...
type Integer struct {
	Value int
}

// Accept ..
func (i *Integer) Accept(visitor interfaces.Visitor) string {
	return visitor.Visit(i)
}

// InitializeInteger ...
func InitializeInteger(value int) *Integer {
	var integer = new(Integer)

	integer.Value = value

	return integer
}
