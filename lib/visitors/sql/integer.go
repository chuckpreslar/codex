package sql

import (
	"fmt"
)

import (
	"github.com/chuckpreslar/codex/lib/interfaces"
	"github.com/chuckpreslar/codex/lib/nodes/base"
)

// Integer ...
type Integer struct{}

// VisitIntegerNode ...
func (i Integer) VisitIntegerNode(node *base.Integer, visitor interfaces.Visitor) string {
	return fmt.Sprintf("%d", node.Value)
}
