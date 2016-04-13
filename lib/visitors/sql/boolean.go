package sql

import (
	"github.com/chuckpreslar/codex/lib/interfaces"
	"github.com/chuckpreslar/codex/lib/nodes/base"
)

// Boolean ...
type Boolean struct{}

// VisitBooleanNode ...
func (b Boolean) VisitBooleanNode(node *base.Boolean, visitor interfaces.Visitor) string {
	if node.Value {
		return "'t'"
	}

	return "'f'"
}
