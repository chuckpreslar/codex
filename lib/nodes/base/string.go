package base

import (
	"github.com/chuckpreslar/codex/lib/interfaces"
)

// String ...
type String struct {
	Value string
}

// Accept ..
func (s *String) Accept(visitor interfaces.Visitor) string {
	return visitor.Visit(s)
}

// InitializeString ...
func InitializeString(value string) *String {
	var str = new(String)

	str.Value = value

	return str
}
